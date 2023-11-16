package entities

import "github.com/samwestmoreland/gonduit/util"

// Project represents a single Phabricator Project.
type Project struct {
	ID               string             `json:"id"`
	PHID             string             `json:"phid"`
	Name             string             `json:"name"`
	ProfileImagePHID string             `json:"profileImagePHID"`
	Icon             string             `json:"icon"`
	Color            string             `json:"color"`
	Members          []string           `json:"members"`
	Slugs            []string           `json:"slugs"`
	DateCreated      util.UnixTimestamp `json:"dateCreated"`
	DateModified     util.UnixTimestamp `json:"dateModified"`
}

func ProjectNameTransaction(name string) Transaction {
	return Transaction{
		Type:  "name",
		Value: name,
	}
}

func ProjectSlugsTransaction(slugs []string) Transaction {
	return Transaction{
		Type:  "slugs",
		Value: slugs,
	}
}

func ProjectIconTransaction(icon string) Transaction {
	return Transaction{
		Type:  "icon",
		Value: icon,
	}
}

func ProjectColorTransaction(color string) Transaction {
	return Transaction{
		Type:  "color",
		Value: color,
	}
}

func ProjectDescriptionTransaction(description string) Transaction {
	return Transaction{
		Type:  "description",
		Value: description,
	}
}

func ProjectViewPolicyTransaction(viewPolicy string) Transaction {
	return Transaction{
		Type:  "view",
		Value: viewPolicy,
	}
}

func ProjectEditPolicyTransaction(editPolicy string) Transaction {
	return Transaction{
		Type:  "edit",
		Value: editPolicy,
	}
}

func ProjectJoinPolicyTransaction(joinPolicy string) Transaction {
	return Transaction{
		Type:  "join",
		Value: joinPolicy,
	}
}

func ProjectAddMembersTransaction(members []PHID) Transaction {
	return Transaction{
		Type:  "members.add",
		Value: members,
	}
}

func ProjectRemoveMembersTransaction(members []PHID) Transaction {
	return Transaction{
		Type:  "members.remove",
		Value: members,
	}
}

func ProjectSetMembersTransaction(members []PHID) Transaction {
	return Transaction{
		Type:  "members.set",
		Value: members,
	}
}
