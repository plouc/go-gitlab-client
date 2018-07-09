# bash completion for glc                                  -*- shell-script -*-

__glc_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Homebrew on Macs have version 1.3 of bash-completion which doesn't include
# _init_completion. This is a very minimal version of that function.
__glc_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

__glc_index_of_word()
{
    local w word=$1
    shift
    index=0
    for w in "$@"; do
        [[ $w = "$word" ]] && return
        index=$((index+1))
    done
    index=-1
}

__glc_contains_word()
{
    local w word=$1; shift
    for w in "$@"; do
        [[ $w = "$word" ]] && return
    done
    return 1
}

__glc_handle_reply()
{
    __glc_debug "${FUNCNAME[0]}"
    case $cur in
        -*)
            if [[ $(type -t compopt) = "builtin" ]]; then
                compopt -o nospace
            fi
            local allflags
            if [ ${#must_have_one_flag[@]} -ne 0 ]; then
                allflags=("${must_have_one_flag[@]}")
            else
                allflags=("${flags[*]} ${two_word_flags[*]}")
            fi
            COMPREPLY=( $(compgen -W "${allflags[*]}" -- "$cur") )
            if [[ $(type -t compopt) = "builtin" ]]; then
                [[ "${COMPREPLY[0]}" == *= ]] || compopt +o nospace
            fi

            # complete after --flag=abc
            if [[ $cur == *=* ]]; then
                if [[ $(type -t compopt) = "builtin" ]]; then
                    compopt +o nospace
                fi

                local index flag
                flag="${cur%=*}"
                __glc_index_of_word "${flag}" "${flags_with_completion[@]}"
                COMPREPLY=()
                if [[ ${index} -ge 0 ]]; then
                    PREFIX=""
                    cur="${cur#*=}"
                    ${flags_completion[${index}]}
                    if [ -n "${ZSH_VERSION}" ]; then
                        # zsh completion needs --flag= prefix
                        eval "COMPREPLY=( \"\${COMPREPLY[@]/#/${flag}=}\" )"
                    fi
                fi
            fi
            return 0;
            ;;
    esac

    # check if we are handling a flag with special work handling
    local index
    __glc_index_of_word "${prev}" "${flags_with_completion[@]}"
    if [[ ${index} -ge 0 ]]; then
        ${flags_completion[${index}]}
        return
    fi

    # we are parsing a flag and don't have a special handler, no completion
    if [[ ${cur} != "${words[cword]}" ]]; then
        return
    fi

    local completions
    completions=("${commands[@]}")
    if [[ ${#must_have_one_noun[@]} -ne 0 ]]; then
        completions=("${must_have_one_noun[@]}")
    fi
    if [[ ${#must_have_one_flag[@]} -ne 0 ]]; then
        completions+=("${must_have_one_flag[@]}")
    fi
    COMPREPLY=( $(compgen -W "${completions[*]}" -- "$cur") )

    if [[ ${#COMPREPLY[@]} -eq 0 && ${#noun_aliases[@]} -gt 0 && ${#must_have_one_noun[@]} -ne 0 ]]; then
        COMPREPLY=( $(compgen -W "${noun_aliases[*]}" -- "$cur") )
    fi

    if [[ ${#COMPREPLY[@]} -eq 0 ]]; then
        declare -F __custom_func >/dev/null && __custom_func
    fi

    # available in bash-completion >= 2, not always present on macOS
    if declare -F __ltrim_colon_completions >/dev/null; then
        __ltrim_colon_completions "$cur"
    fi

    # If there is only 1 completion and it is a flag with an = it will be completed
    # but we don't want a space after the =
    if [[ "${#COMPREPLY[@]}" -eq "1" ]] && [[ $(type -t compopt) = "builtin" ]] && [[ "${COMPREPLY[0]}" == --*= ]]; then
       compopt -o nospace
    fi
}

# The arguments should be in the form "ext1|ext2|extn"
__glc_handle_filename_extension_flag()
{
    local ext="$1"
    _filedir "@(${ext})"
}

__glc_handle_subdirs_in_dir_flag()
{
    local dir="$1"
    pushd "${dir}" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1
}

__glc_handle_flag()
{
    __glc_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    # if a command required a flag, and we found it, unset must_have_one_flag()
    local flagname=${words[c]}
    local flagvalue
    # if the word contained an =
    if [[ ${words[c]} == *"="* ]]; then
        flagvalue=${flagname#*=} # take in as flagvalue after the =
        flagname=${flagname%=*} # strip everything after the =
        flagname="${flagname}=" # but put the = back
    fi
    __glc_debug "${FUNCNAME[0]}: looking for ${flagname}"
    if __glc_contains_word "${flagname}" "${must_have_one_flag[@]}"; then
        must_have_one_flag=()
    fi

    # if you set a flag which only applies to this command, don't show subcommands
    if __glc_contains_word "${flagname}" "${local_nonpersistent_flags[@]}"; then
      commands=()
    fi

    # keep flag value with flagname as flaghash
    # flaghash variable is an associative array which is only supported in bash > 3.
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        if [ -n "${flagvalue}" ] ; then
            flaghash[${flagname}]=${flagvalue}
        elif [ -n "${words[ $((c+1)) ]}" ] ; then
            flaghash[${flagname}]=${words[ $((c+1)) ]}
        else
            flaghash[${flagname}]="true" # pad "true" for bool flag
        fi
    fi

    # skip the argument to a two word flag
    if __glc_contains_word "${words[c]}" "${two_word_flags[@]}"; then
        c=$((c+1))
        # if we are looking for a flags value, don't show commands
        if [[ $c -eq $cword ]]; then
            commands=()
        fi
    fi

    c=$((c+1))

}

__glc_handle_noun()
{
    __glc_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    if __glc_contains_word "${words[c]}" "${must_have_one_noun[@]}"; then
        must_have_one_noun=()
    elif __glc_contains_word "${words[c]}" "${noun_aliases[@]}"; then
        must_have_one_noun=()
    fi

    nouns+=("${words[c]}")
    c=$((c+1))
}

__glc_handle_command()
{
    __glc_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    local next_command
    if [[ -n ${last_command} ]]; then
        next_command="_${last_command}_${words[c]//:/__}"
    else
        if [[ $c -eq 0 ]]; then
            next_command="_glc_root_command"
        else
            next_command="_${words[c]//:/__}"
        fi
    fi
    c=$((c+1))
    __glc_debug "${FUNCNAME[0]}: looking for ${next_command}"
    declare -F "$next_command" >/dev/null && $next_command
}

__glc_handle_word()
{
    if [[ $c -ge $cword ]]; then
        __glc_handle_reply
        return
    fi
    __glc_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"
    if [[ "${words[c]}" == -* ]]; then
        __glc_handle_flag
    elif __glc_contains_word "${words[c]}" "${commands[@]}"; then
        __glc_handle_command
    elif [[ $c -eq 0 ]]; then
        __glc_handle_command
    elif __glc_contains_word "${words[c]}" "${command_aliases[@]}"; then
        # aliashash variable is an associative array which is only supported in bash > 3.
        if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
            words[c]=${aliashash[${words[c]}]}
            __glc_handle_command
        else
            __glc_handle_noun
        fi
    else
        __glc_handle_noun
    fi
    __glc_handle_word
}

_glc_add_alias()
{
    last_command="glc_add_alias"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_group()
{
    last_command="glc_add_group"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_group-var()
{
    last_command="glc_add_group-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project()
{
    last_command="glc_add_project"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-badge()
{
    last_command="glc_add_project-badge"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-branch()
{
    last_command="glc_add_project-branch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--branch=")
    two_word_flags+=("-b")
    local_nonpersistent_flags+=("--branch=")
    flags+=("--ref=")
    two_word_flags+=("-r")
    local_nonpersistent_flags+=("--ref=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-hook()
{
    last_command="glc_add_project-hook"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-protected-branch()
{
    last_command="glc_add_project-protected-branch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-star()
{
    last_command="glc_add_project-star"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add_project-var()
{
    last_command="glc_add_project-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_add()
{
    last_command="glc_add"

    command_aliases=()

    commands=()
    commands+=("alias")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("a")
        aliashash["a"]="alias"
    fi
    commands+=("group")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("g")
        aliashash["g"]="group"
    fi
    commands+=("group-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("gv")
        aliashash["gv"]="group-var"
    fi
    commands+=("project")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("p")
        aliashash["p"]="project"
    fi
    commands+=("project-badge")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pbdg")
        aliashash["pbdg"]="project-badge"
    fi
    commands+=("project-branch")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pb")
        aliashash["pb"]="project-branch"
    fi
    commands+=("project-hook")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ph")
        aliashash["ph"]="project-hook"
    fi
    commands+=("project-protected-branch")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ppb")
        aliashash["ppb"]="project-protected-branch"
    fi
    commands+=("project-star")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ps")
        aliashash["ps"]="project-star"
    fi
    commands+=("project-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pv")
        aliashash["pv"]="project-var"
    fi

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_current-user()
{
    last_command="glc_get_current-user"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_group()
{
    last_command="glc_get_group"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--with-custom-attributes")
    flags+=("-x")
    local_nonpersistent_flags+=("--with-custom-attributes")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_group-var()
{
    last_command="glc_get_group-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_namespace()
{
    last_command="glc_get_namespace"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project()
{
    last_command="glc_get_project"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--statistics")
    flags+=("-s")
    local_nonpersistent_flags+=("--statistics")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project-badge()
{
    last_command="glc_get_project-badge"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project-branch()
{
    last_command="glc_get_project-branch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project-hook()
{
    last_command="glc_get_project-hook"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project-pipeline()
{
    last_command="glc_get_project-pipeline"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_project-var()
{
    last_command="glc_get_project-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_runner()
{
    last_command="glc_get_runner"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get_user()
{
    last_command="glc_get_user"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_get()
{
    last_command="glc_get"

    command_aliases=()

    commands=()
    commands+=("current-user")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("cu")
        aliashash["cu"]="current-user"
    fi
    commands+=("group")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("g")
        aliashash["g"]="group"
    fi
    commands+=("group-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("gv")
        aliashash["gv"]="group-var"
    fi
    commands+=("namespace")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ns")
        aliashash["ns"]="namespace"
    fi
    commands+=("project")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("p")
        aliashash["p"]="project"
    fi
    commands+=("project-badge")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pbdg")
        aliashash["pbdg"]="project-badge"
    fi
    commands+=("project-branch")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pb")
        aliashash["pb"]="project-branch"
    fi
    commands+=("project-hook")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ph")
        aliashash["ph"]="project-hook"
    fi
    commands+=("project-pipeline")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pp")
        aliashash["pp"]="project-pipeline"
    fi
    commands+=("project-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pv")
        aliashash["pv"]="project-var"
    fi
    commands+=("runner")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("r")
        aliashash["r"]="runner"
    fi
    commands+=("user")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("u")
        aliashash["u"]="user"
    fi

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_init()
{
    last_command="glc_init"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_aliases()
{
    last_command="glc_ls_aliases"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_group-vars()
{
    last_command="glc_ls_group-vars"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_groups()
{
    last_command="glc_ls_groups"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--all")
    local_nonpersistent_flags+=("--all")
    flags+=("--owned")
    local_nonpersistent_flags+=("--owned")
    flags+=("--search=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--search=")
    flags+=("--statistics")
    local_nonpersistent_flags+=("--statistics")
    flags+=("--with-custom-attributes")
    flags+=("-x")
    local_nonpersistent_flags+=("--with-custom-attributes")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_namespaces()
{
    last_command="glc_ls_namespaces"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--search=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--search=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-badges()
{
    last_command="glc_ls_project-badges"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-branches()
{
    last_command="glc_ls_project-branches"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--search=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--search=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-hooks()
{
    last_command="glc_ls_project-hooks"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-members()
{
    last_command="glc_ls_project-members"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--query=")
    two_word_flags+=("-q")
    local_nonpersistent_flags+=("--query=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-pipelines()
{
    last_command="glc_ls_project-pipelines"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-protected-branches()
{
    last_command="glc_ls_project-protected-branches"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_project-vars()
{
    last_command="glc_ls_project-vars"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_projects()
{
    last_command="glc_ls_projects"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--archived")
    local_nonpersistent_flags+=("--archived")
    flags+=("--membership")
    local_nonpersistent_flags+=("--membership")
    flags+=("--owned")
    local_nonpersistent_flags+=("--owned")
    flags+=("--search=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--search=")
    flags+=("--starred")
    local_nonpersistent_flags+=("--starred")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_runners()
{
    last_command="glc_ls_runners"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--all")
    local_nonpersistent_flags+=("--all")
    flags+=("--scope=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--scope=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls_users()
{
    last_command="glc_ls_users"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--active")
    local_nonpersistent_flags+=("--active")
    flags+=("--blocked")
    local_nonpersistent_flags+=("--blocked")
    flags+=("--search=")
    two_word_flags+=("-s")
    local_nonpersistent_flags+=("--search=")
    flags+=("--username=")
    two_word_flags+=("-u")
    local_nonpersistent_flags+=("--username=")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_ls()
{
    last_command="glc_ls"

    command_aliases=()

    commands=()
    commands+=("aliases")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("a")
        aliashash["a"]="aliases"
    fi
    commands+=("group-vars")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("gv")
        aliashash["gv"]="group-vars"
    fi
    commands+=("groups")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("g")
        aliashash["g"]="groups"
    fi
    commands+=("namespaces")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ns")
        aliashash["ns"]="namespaces"
    fi
    commands+=("project-badges")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pbdg")
        aliashash["pbdg"]="project-badges"
    fi
    commands+=("project-branches")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pb")
        aliashash["pb"]="project-branches"
    fi
    commands+=("project-hooks")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ph")
        aliashash["ph"]="project-hooks"
    fi
    commands+=("project-members")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pm")
        aliashash["pm"]="project-members"
    fi
    commands+=("project-pipelines")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pp")
        aliashash["pp"]="project-pipelines"
    fi
    commands+=("project-protected-branches")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ppb")
        aliashash["ppb"]="project-protected-branches"
    fi
    commands+=("project-vars")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pv")
        aliashash["pv"]="project-vars"
    fi
    commands+=("projects")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("p")
        aliashash["p"]="projects"
    fi
    commands+=("runners")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("r")
        aliashash["r"]="runners"
    fi
    commands+=("users")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("u")
        aliashash["u"]="users"
    fi

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--page=")
    two_word_flags+=("-p")
    flags+=("--per_page=")
    two_word_flags+=("-l")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_alias()
{
    last_command="glc_rm_alias"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_group()
{
    last_command="glc_rm_group"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_group-var()
{
    last_command="glc_rm_group-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project()
{
    last_command="glc_rm_project"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-badge()
{
    last_command="glc_rm_project-badge"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-branch()
{
    last_command="glc_rm_project-branch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-hook()
{
    last_command="glc_rm_project-hook"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-merged-branches()
{
    last_command="glc_rm_project-merged-branches"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-protected-branch()
{
    last_command="glc_rm_project-protected-branch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-star()
{
    last_command="glc_rm_project-star"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm_project-var()
{
    last_command="glc_rm_project-var"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")
    flags+=("--yes")
    flags+=("-y")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_rm()
{
    last_command="glc_rm"

    command_aliases=()

    commands=()
    commands+=("alias")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("a")
        aliashash["a"]="alias"
    fi
    commands+=("group")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("g")
        aliashash["g"]="group"
    fi
    commands+=("group-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("gv")
        aliashash["gv"]="group-var"
    fi
    commands+=("project")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("p")
        aliashash["p"]="project"
    fi
    commands+=("project-badge")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pbdg")
        aliashash["pbdg"]="project-badge"
    fi
    commands+=("project-branch")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pb")
        aliashash["pb"]="project-branch"
    fi
    commands+=("project-hook")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ph")
        aliashash["ph"]="project-hook"
    fi
    commands+=("project-merged-branches")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pmb")
        aliashash["pmb"]="project-merged-branches"
    fi
    commands+=("project-protected-branch")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ppb")
        aliashash["ppb"]="project-protected-branch"
    fi
    commands+=("project-star")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("ps")
        aliashash["ps"]="project-star"
    fi
    commands+=("project-var")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("pv")
        aliashash["pv"]="project-var"
    fi

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--yes")
    flags+=("-y")
    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_version()
{
    last_command="glc_version"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_glc_root_command()
{
    last_command="glc"

    command_aliases=()

    commands=()
    commands+=("add")
    commands+=("get")
    commands+=("init")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("i")
        aliashash["i"]="init"
    fi
    commands+=("ls")
    commands+=("rm")
    commands+=("version")
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        command_aliases+=("v")
        aliashash["v"]="version"
    fi

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--alias=")
    two_word_flags+=("-a")
    flags+=("--interactive")
    flags+=("-i")
    flags+=("--no-color")
    flags+=("--output-destination=")
    two_word_flags+=("-o")
    flags+=("--output-format=")
    two_word_flags+=("-f")
    flags+=("--silent")
    flags+=("--verbose")
    flags+=("-v")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

__start_glc()
{
    local cur prev words cword
    declare -A flaghash 2>/dev/null || :
    declare -A aliashash 2>/dev/null || :
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -s || return
    else
        __glc_init_completion -n "=" || return
    fi

    local c=0
    local flags=()
    local two_word_flags=()
    local local_nonpersistent_flags=()
    local flags_with_completion=()
    local flags_completion=()
    local commands=("glc")
    local must_have_one_flag=()
    local must_have_one_noun=()
    local last_command
    local nouns=()

    __glc_handle_word
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_glc glc
else
    complete -o default -o nospace -F __start_glc glc
fi

# ex: ts=4 sw=4 et filetype=sh
