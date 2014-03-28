package gogitlab

import (
	"encoding/json"
)

const (
	// ID
	user_keys        = "/user/keys"     // Get current user keys
	user_key         = "/user/keys/:id" // Get user key by id
	custom_user_keys = "/user/:id/keys" // Create key for user with :id
)

func (g *Gitlab) UserKeys() ([]*UserKey, error) {

	url := g.ResourceUrl(user_keys, nil)

	var keys []*UserKey

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &keys)
	}

	return keys, err
}

// /*
// Get single project deploy key.

//     GET /projects/:id/keys/:key_id

// Parameters:

//     id     The ID of a project
//     key_id The ID of a key

// */
// func (g *Gitlab) UserKey(id string) (*UserKey, error) {

// 	url := g.ResourceUrl(user_key, map[string]string{":id": id})

// 	var key *UserKey

// 	contents, err := g.buildAndExecRequest("GET", url, nil)
// 	if err == nil {
// 		err = json.Unmarshal(contents, &Key)
// 	}

// 	return key, err
// }

// /*
// Add deploy key to project.

//     POST /projects/:id/keys

// Parameters:

//     id    The ID of a project
//     title The key title
//     key   The key value

// */
// func (g *Gitlab) AddKey(title, key string) error {

// 	path := g.ResourceUrl(user_keys, nil)

// 	var err error

// 	v := url.Values{}
// 	v.Set("title", title)
// 	v.Set("key", key)

// 	body := v.Encode()

// 	_, err = g.buildAndExecRequest("POST", path, []byte(body))

// 	return err
// }

// /*
// Remove deploy key from project

//     DELETE /projects/:id/keys/:key_id

// Parameters:

//     id     The ID of a project
//     key_id The ID of a key

// */
// func (g *Gitlab) RemoveKey(id string) error {

// 	url := g.ResourceUrl(user_key, map[string]string{":id": id})

// 	var err error

// 	_, err = g.buildAndExecRequest("DELETE", url, nil)

// 	return err
// }
