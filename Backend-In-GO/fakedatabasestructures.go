package main

type (
	usernamePassword struct {
		username string
		password string
		personalScore int
	}

	hallOfFameEntry struct {
		user *usernamePassword
		score int
	}

	binaryNode struct {
		user *usernamePassword
		left  *binaryNode // pointer to point to left node
		right *binaryNode // pointer to point to right node
	}
	
	bst struct {
		root *binaryNode
		size int
	}
)

// thread safe bst as a sub for map lol, storing all user info, sorted by username a-z A-Z
//create
func (bst *bst) createUser(username string, password string, personalScore int) error {
	if bst.size == 0 {
		bst.root = &binaryNode{
			user : &usernamePassword{
				username: username,
				password: password,
				personalScore: personalScore,
			},
			left : nil,
			right : nil,
		}

		bst.size++
	}

	if username == "" || password == 

	return nil
}

//read (no need lock)
//for bst to make sure username not already taken

//for when user login, retrieve account info

//update

//delete

// thread safe prio queue for daily score

// thread safe prio queue for hall of fame