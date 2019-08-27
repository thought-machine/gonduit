package entities

import (
	"fmt"

	"github.com/thought-machine/gonduit/util"
)

// ManiphestTask represents a single task on Maniphest.
type ManiphestTask struct {
	ID                 string             `json:"id"`
	PHID               string             `json:"phid"`
	AuthorPHID         string             `json:"authorPHID"`
	OwnerPHID          string             `json:"ownerPHID"`
	CCPHIDs            []string           `json:"ccPHIDs"`
	Status             string             `json:"status"`
	StatusName         string             `json:"statusName"`
	IsClosed           bool               `json:"isClosed"`
	Priority           string             `json:"priority"`
	PriorityColor      string             `json:"priorityColor"`
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	ProjectPHIDs       []string           `json:"projectPHIDs"`
	URI                string             `json:"uri"`
	Auxiliary          interface{}        `json:"auxiliary"`
	ObjectName         string             `json:"objectName"`
	DateCreated        util.UnixTimestamp `json:"dateCreated"`
	DateModified       util.UnixTimestamp `json:"dateModified"`
	DependsOnTaskPHIDs []string           `json:"dependsOnTaskPHIDs"`
}

type ManiphestPoints int
type ManiphestPriority string
type ManiphestStatus string
type ManiphestSubtype string

// ManiphestParentTransaction marks this task as a subtask of the given task.
func ManiphestParentTransaction(phid PHID) Transaction {
	return Transaction{
		Type:  "parent",
		Value: phid,
	}
}

// ManiphestColumnTransaction moves the task into the given columns. We don't support advanced usage, if you require
// the ability to move the task before/after others in the column, construct the transaction yourself.
func ManiphestColumnTransaction(columnPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "column",
		Value: columnPHIDs,
	}
}

// ManiphestSpaceTransaction moves the task between Spaces.
func ManiphestSpaceTransaction(spacePHID PHID) Transaction {
	return Transaction{
		Type:  "space",
		Value: spacePHID,
	}
}

// ManiphestTitleTransaction renames the task.
func ManiphestTitleTransaction(title string) Transaction {
	return Transaction{
		Type:  "title",
		Value: title,
	}
}

// ManiphestOwnerTransaction reassigns the task.
func ManiphestOwnerTransaction(owner PHID) Transaction {
	return Transaction{
		Type:  "owner",
		Value: owner,
	}
}

// ManiphestStatusTransaction changes the task status. Check your install for valid argument values.
func ManiphestStatusTransaction(status ManiphestStatus) Transaction {
	return Transaction{
		Type:  "status",
		Value: status,
	}
}

// ManiphestPriorityTransaction changes the priority of the task. Check your install for valid argument values.
func ManiphestPriorityTransaction(priority ManiphestPriority) Transaction {
	return Transaction{
		Type:  "priority",
		Value: priority,
	}
}

// ManiphestPointsTransaction changes the task point value. If you require decimal points, create the transaction yourself.
func ManiphestPointsTransaction(points ManiphestPoints) Transaction {
	return Transaction{
		Type:  "points",
		Value: points,
	}
}

// ManiphestDescriptionTransaction updates the task description.
func ManiphestDescriptionTransaction(description string) Transaction {
	return Transaction{
		Type:  "description",
		Value: description,
	}
}

// ManiphestAddParentsTransaction changes the parents of this task.
func ManiphestAddParentsTransaction(parentPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "parents.add",
		Value: parentPHIDs,
	}
}

// ManiphestRemoveParentsTransaction changes the parents of this task.
func ManiphestRemoveParentsTransaction(parentPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "parents.remove",
		Value: parentPHIDs,
	}
}

// ManiphestSetParentsTransaction changes the parents of this task.
func ManiphestSetParentsTransaction(parentPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "parents.set",
		Value: parentPHIDs,
	}
}

// ManiphestAddSubtasksTransaction changes the subtasks of this task.
func ManiphestAddSubtasksTransaction(subtaskPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subtasks.add",
		Value: subtaskPHIDs,
	}
}

// ManiphestRemoveSubtasksTransaction changes the subtasks of this task.
func ManiphestRemoveSubtasksTransaction(subtaskPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subtasks.remove",
		Value: subtaskPHIDs,
	}
}

// ManiphestSetSubtasksTransaction changes the subtasks of this task.
func ManiphestSetSubtasksTransaction(subtaskPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subtasks.set",
		Value: subtaskPHIDs,
	}
}

// ManiphestViewPolicyTransaction changes the view policy of the object.
func ManiphestViewPolicyTransaction(policy string) Transaction {
	return Transaction{
		Type:  "view",
		Value: policy,
	}
}

// ManiphestEditPolicyTransaction changes the edit policy of the object.
func ManiphestEditPolicyTransaction(policy string) Transaction {
	return Transaction{
		Type:  "edit",
		Value: policy,
	}
}

// ManiphestAddProjectsTransaction adds project tags.
func ManiphestAddProjectsTransaction(projectPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "projects.add",
		Value: projectPHIDs,
	}
}

// ManiphestRemoveProjectsTransaction removes project tags.
func ManiphestRemoveProjectsTransaction(projectPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "projects.remove",
		Value: projectPHIDs,
	}
}

// ManiphestSetProjectsTransaction sets project tags, overwriting current value.
func ManiphestSetProjectsTransaction(projectPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "projects.set",
		Value: projectPHIDs,
	}
}

// ManiphestAddSubscribersTransaction adds subscribers.
func ManiphestAddSubscribersTransaction(userPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subscribers.add",
		Value: userPHIDs,
	}
}

// ManiphestRemoveSubscribersTransaction removes subscribers.
func ManiphestRemoveSubscribersTransaction(userPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subscribers.remove",
		Value: userPHIDs,
	}
}

// ManiphestSetSubscribersTransaction sets subscribers, overwriting current value.
func ManiphestSetSubscribersTransaction(userPHIDs ...PHID) Transaction {
	return Transaction{
		Type:  "subscribers.set",
		Value: userPHIDs,
	}
}

// ManiphestSubtypeTransaction changes the object subtype.
func ManiphestSubtypeTransaction(subtype ManiphestSubtype) Transaction {
	return Transaction{
		Type:  "subtype",
		Value: subtype,
	}
}

// ManiphestCommentTransaction makes comments.
func ManiphestCommentTransaction(comment string) Transaction {
	return Transaction{
		Type:  "comment",
		Value: comment,
	}
}

// ManiphestMFASignTransaction signs this transaction group with MFA.
func ManiphestMFASignTransaction(sign bool) Transaction {
	return Transaction{
		Type:  "mfa",
		Value: sign,
	}
}

// ManiphestCustomFieldTransaction allows the caller to set the value of the given field.
func ManiphestCustomFieldTransaction(fieldName string, value string) Transaction {
	return Transaction{
		Type:  fmt.Sprintf("custom.%s", fieldName),
		Value: value,
	}
}
