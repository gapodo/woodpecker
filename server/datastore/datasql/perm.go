package datasql

import (
	"github.com/drone/drone/shared/model"
	"github.com/russross/meddler"
)

type Permstore struct {
	meddler.DB
}

func NewPermstore(db meddler.DB) *Permstore {
	return &Permstore{db}
}

// GetPerm retrieves the User's permission from
// the datastore for the given repository.
func (db *Repostore) GetPerm(user *model.User, repo *model.Repo) (*model.Perm, error) {
	var perm = new(model.Perm)
	var err = meddler.QueryRow(db, perm, permQuery, user.ID, repo.ID)
	return perm, err
}

// PostPerm saves permission in the datastore.
func (db *Repostore) PostPerm(perm *model.Perm) error {
	return meddler.Save(db, permTable, perm)
}

// PutPerm saves permission in the datastore.
func (db *Repostore) PutPerm(perm *model.Perm) error {
	return meddler.Save(db, permTable, perm)
}

// DelPerm removes permission from the datastore.
func (db *Repostore) DelPerm(perm *model.Perm) error {
	var _, err = db.Exec(permDeleteStmt, perm.ID)
	return err
}

// Permission table name in database.
const permTable = "perms"

// SQL query to retrieve a user's permission to
// access a repository.
const permQuery = `
SELECT *
FROM perms
WHERE user_id=?
AND   repo_id=?
LIMIT 1
`

// SQL statement to delete a User by ID.
const permDeleteStmt = `
DELETE FROM perms
WHERE perm_id=?
`
