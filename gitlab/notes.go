package gitlab

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

const (
	ProjectIssueNotesApiPath        = "/projects/:id/issues/:issue_iid/notes"
	ProjectIssueNoteApiPath         = "/projects/:id/issues/:issue_iid/notes/:note_id"
	ProjectSnippetNotesApiPath      = "/projects/:id/snippets/:snippet_id/notes"
	ProjectSnippetNoteApiPath       = "/projects/:id/snippets/:snippet_id/notes/:note_id"
	ProjectMergeRequestNotesApiPath = "/projects/:id/merge_requests/:merge_request_iid/notes"
	ProjectMergeRequestNoteApiPath  = "/projects/:id/merge_requests/:merge_request_iid/notes/:note_id"
	GroupEpicNotesApiPath           = "/groups/:id/epics/:epic_id/notes"
	GroupEpicNoteApiPath            = "/groups/:id/epics/:epic_id/notes/:note_id"
)

type Note struct {
	Id           int    `json:"id"            yaml:"id"`
	Body         string `json:"body"          yaml:"body"`
	Attachment   string `json:"attachment"    yaml:"attachment"`
	Title        string `json:"title"         yaml:"title"`
	FileName     string `json:"file_name"     yaml:"file_name"`
	CreatedAtRaw string `json:"created_at"    yaml:"created_at"`
	UpdatedAtRaw string `json:"updated_at"    yaml:"updated_at"`
	ExpiresAtRaw string `json:"expires_at"    yaml:"expires_at"`
	System       bool   `json:"system"        yaml:"system"`
	Resolvable   bool   `json:"resolvable"    yaml:"resolvable"`
	NoteableId   int    `json:"noteable_id"   yaml:"noteable_id"`
	NoteableIid  int    `json:"noteable_iid"  yaml:"noteable_iid"`
	NoteableType string `json:"noteable_type" yaml:"noteable_type"`
	Author       struct {
		Id           int    `json:"id"         yaml:"id"`
		Username     string `json:"username"   yaml:"username"`
		Email        string `json:"email"      yaml:"email"`
		Name         string `json:"name"       yaml:"name"`
		State        string `json:"state"      yaml:"state"`
		CreatedAtRaw string `json:"created_at" yaml:"created_at"`
		AvatarUrm    string `json:"avatar_url" yaml:"avatar_url"`
		WebUrl       string `json:"web_url"    yaml:"web_url"`
	} `json:"author" yaml:"author"`
}

type NoteAddPayload struct {
	Body string `json:"body"`
}

func (n *Note) RenderJson(w io.Writer) error {
	return renderJson(w, n)
}

func (n *Note) RenderYaml(w io.Writer) error {
	return renderYaml(w, n)
}

type NoteCollection struct {
	Items []*Note
}

func (c *NoteCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *NoteCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

type NotesOptions struct {
	PaginationOptions
	SortOptions
}

func (g *Gitlab) getNotes(u *url.URL) (*NoteCollection, *ResponseMeta, error) {
	collection := new(NoteCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectIssueNotes(projectId string, issueIid int, o *NotesOptions) (*NoteCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectIssueNotesApiPath, map[string]string{
		":id":        projectId,
		":issue_iid": strconv.Itoa(issueIid),
	}, o)

	return g.getNotes(u)
}

func (g *Gitlab) ProjectSnippetNotes(projectId string, snippetId int, o *NotesOptions) (*NoteCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectSnippetNotesApiPath, map[string]string{
		":id":         projectId,
		":snippet_id": strconv.Itoa(snippetId),
	}, o)

	return g.getNotes(u)
}

func (g *Gitlab) ProjectMergeRequestNotes(projectId string, mergeRequestIid int, o *NotesOptions) (*NoteCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(ProjectMergeRequestNotesApiPath, map[string]string{
		":id":                projectId,
		":merge_request_iid": strconv.Itoa(mergeRequestIid),
	}, o)

	return g.getNotes(u)
}

func (g *Gitlab) GroupEpicNotes(groupId string, epicId int, o *NotesOptions) (*NoteCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(GroupEpicNotesApiPath, map[string]string{
		":id":      groupId,
		":epic_id": strconv.Itoa(epicId),
	}, o)

	return g.getNotes(u)
}

func (g *Gitlab) getNote(u *url.URL) (*Note, *ResponseMeta, error) {
	note := new(Note)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &note)
	}

	return note, meta, err
}

func (g *Gitlab) ProjectIssueNote(projectId string, issueIid, noteId int) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectIssueNoteApiPath, map[string]string{
		":id":        projectId,
		":issue_iid": strconv.Itoa(issueIid),
		":note_id":   strconv.Itoa(noteId),
	})

	return g.getNote(u)
}

func (g *Gitlab) ProjectSnippetNote(projectId string, snippetId, noteId int) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectSnippetNoteApiPath, map[string]string{
		":id":         projectId,
		":snippet_id": strconv.Itoa(snippetId),
		":note_id":    strconv.Itoa(noteId),
	})

	return g.getNote(u)
}

func (g *Gitlab) ProjectMergeRequestNote(projectId string, mergeRequestIid, noteId int) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectMergeRequestNoteApiPath, map[string]string{
		":id":                projectId,
		":merge_request_iid": strconv.Itoa(mergeRequestIid),
		":note_id":           strconv.Itoa(noteId),
	})

	return g.getNote(u)
}

func (g *Gitlab) GroupEpicNote(groupId string, epicId, noteId int) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(GroupEpicNoteApiPath, map[string]string{
		":id":      groupId,
		":epic_id": strconv.Itoa(epicId),
		":note_id": strconv.Itoa(noteId),
	})

	return g.getNote(u)
}

func (g *Gitlab) removeNote(u *url.URL) (*ResponseMeta, error) {
	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}

func (g *Gitlab) RemoveProjectIssueNote(projectId string, issueIid, noteId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectIssueNoteApiPath, map[string]string{
		":id":        projectId,
		":issue_iid": strconv.Itoa(issueIid),
		":note_id":   strconv.Itoa(noteId),
	})

	return g.removeNote(u)
}

func (g *Gitlab) RemoveProjectSnippetNote(projectId string, snippetId, noteId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectSnippetNoteApiPath, map[string]string{
		":id":         projectId,
		":snippet_id": strconv.Itoa(snippetId),
		":note_id":    strconv.Itoa(noteId),
	})

	return g.removeNote(u)
}

func (g *Gitlab) RemoveProjectMergeRequestNote(projectId string, mergeRequestIid, noteId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(ProjectMergeRequestNoteApiPath, map[string]string{
		":id":                projectId,
		":merge_request_iid": strconv.Itoa(mergeRequestIid),
		":note_id":           strconv.Itoa(noteId),
	})

	return g.removeNote(u)
}

func (g *Gitlab) RemoveGroupEpicNote(groupId string, epicId, noteId int) (*ResponseMeta, error) {
	u := g.ResourceUrl(GroupEpicNoteApiPath, map[string]string{
		":id":      groupId,
		":epic_id": strconv.Itoa(epicId),
		":note_id": strconv.Itoa(noteId),
	})

	return g.removeNote(u)
}

func (g *Gitlab) addNote(u *url.URL, note *NoteAddPayload) (*Note, *ResponseMeta, error) {
	noteJson, err := json.Marshal(note)
	if err != nil {
		return nil, nil, err
	}

	var createdNote *Note
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), noteJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdNote)
	}

	return createdNote, meta, err
}

func (g *Gitlab) AddProjectIssueNote(projectId string, issueIid int, note *NoteAddPayload) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectIssueNotesApiPath, map[string]string{
		":id":        projectId,
		":issue_iid": strconv.Itoa(issueIid),
	})

	return g.addNote(u, note)
}

func (g *Gitlab) AddProjectSnippetNote(projectId string, snippetId int, note *NoteAddPayload) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectSnippetNotesApiPath, map[string]string{
		":id":         projectId,
		":snippet_id": strconv.Itoa(snippetId),
	})

	return g.addNote(u, note)
}

func (g *Gitlab) AddProjectMergeRequestNote(projectId string, mergeRequestIid int, note *NoteAddPayload) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectMergeRequestNotesApiPath, map[string]string{
		":id":                projectId,
		":merge_request_iid": strconv.Itoa(mergeRequestIid),
	})

	return g.addNote(u, note)
}

func (g *Gitlab) AddGroupEpicNote(groupId string, epicId int, note *NoteAddPayload) (*Note, *ResponseMeta, error) {
	u := g.ResourceUrl(GroupEpicNotesApiPath, map[string]string{
		":id":      groupId,
		":epic_id": strconv.Itoa(epicId),
	})

	return g.addNote(u, note)
}
