package main

import "errors"

type (
	userPasswordScore struct {
		username string
		password string
		personalScore int
	}

	queueRecord struct {
		username string
		score int
	}

	prioQueue struct {
		//front is highest score
		//back is lowest score
		front *qNode
		back *qNode
		size int
	}

	qNode struct {
		item *queueRecord
		next *qNode
	}

	binaryNode struct {
		user *userPasswordScore
		left  *binaryNode // pointer to point to left node
		right *binaryNode // pointer to point to right node
	}
	
	bst struct {
		root *binaryNode
		size int
	}
)

// thread safe bst as a sub for map lol, storing all user info, sorted by username a-z A-Z
// returns error if username taken
func (bst *bst) createUser(username string, password string, personalScore int) error {
	mu.Lock()
	//if empty db just add
	if bst.size == 0 {
		bst.root = &binaryNode{
			user : &userPasswordScore{
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
				user : &userPasswordScore{
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

//pass in root of the bst, returns found user node if avail
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

//update

//delete

//internal helper to add scores
func (p *prioQueue) enqueue(username string, score int) error {
	newNode := &qNode{
		item: &queueRecord{
			username: username,
			score: score,
		},
		next: nil,
	}

	mu.Lock()

	if p.front == nil {
		p.front = newNode
		p.back = newNode

	} else {
		if newNode.item.score > p.front.item.score {
			newNode.next = p.front
			p.front = newNode
		} else {
			current := p.front
			for current.next != nil && current.next.item.score > newNode.item.score {
                current = current.next
            }
			newNode.next = current.next
			current.next = newNode

			if newNode.next == nil {
				p.back = newNode
			}
		}

	}
	p.size++
	mu.Unlock()
	return nil
}

//returns an error if score not entered into queue (queue was full AND score wasnt high enough for top 10)
func (p *prioQueue) enterNewScoreIntoQueue(username string, score int) error {
	mu.Lock()

	if p.size < 10 {
		//HAVE TO RELEASE LOCK BEFORE GOING INTO ANOTHER LOCK
		mu.Unlock()
		p.enqueue(username, score)
		return nil
	} else if score > p.back.item.score{
		p.back = nil
		p.size--
		mu.Unlock()
		p.enqueue(username,score)
		return nil
	}

	mu.Unlock()
	return errors.New("Score was not high enough to enter")
}

func (p *prioQueue) isEmpty() bool {
	return p.size == 0
}