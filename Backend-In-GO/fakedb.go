package main

import "errors"

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

	prioQueue struct {

	}
)

// thread safe bst as a sub for map lol, storing all user info, sorted by username a-z A-Z
// returns error if username taken
func (bst *bst) createUser(username string, password string, personalScore int) error {
	mu.Lock()
	//if empty db just add
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
	} else {
		//check if username taken?
		_, err := bst.findUserNode(bst.root, username)
		//cus return err if user does not exist
		if err != nil {
			bst.insertNodeHelper(&bst.root, &binaryNode{
				user : &usernamePassword{
					username: username,
					password: password,
					personalScore: personalScore,
				},
				left : nil,
				right : nil,
			})
		} else {
			//username taken
			mu.Unlock()
			return errors.New("Username is taken")
		}
	}

	bst.size++
	mu.Unlock()
	return nil
}

//im not balancing the tree, aint no time to learn and apply everything on top of other stuff to implement in a day
//if it becomes a linkedlist, it is what it is
func (bst *bst) insertNodeHelper(root **binaryNode, userNode *binaryNode) {
	if *root == nil {
		*root = userNode
		return
	}

	if userNode.user.username < (*root).user.username {
		bst.insertNodeHelper(&((*root).left), userNode)
	} else {
		bst.insertNodeHelper(&((*root).right), userNode)
	}
}

//read (no need lock)
func (bst *bst) findUserNode(root *binaryNode, username string) (*binaryNode, error) {
	if root == nil {
		return nil, errors.New("User does not exist")
	}

	if username == root.user.username {
		return root, nil
	}

	if username > root.user.username {
		return bst.findUserNode(root.right, username)
	} else {
		return bst.findUserNode(root.left, username)
	}
}

//for when user login, retrieve account info

//update

//delete

// thread safe prio queue for daily score

// thread safe prio queue for hall of fame