package gonduit

import (
	"github.com/samwestmoreland/gonduit/requests"
	"github.com/samwestmoreland/gonduit/responses"
)

// DiffusionQueryCommits performs a call to diffusion.querycommits.
// Deprecated: This method is frozen and will eventually be deprecated. New code should use "diffusion.commit.search" instead.
func (c *Conn) DiffusionQueryCommits(
	req requests.DiffusionQueryCommitsRequest,
) (*responses.DiffusionQueryCommitsResponse, error) {
	var res responses.DiffusionQueryCommitsResponse

	if err := c.Call("diffusion.querycommits", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DiffusionCommitSearch performs a call to diffusion.commit.search
func (c *Conn) DiffusionCommitSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("diffusion.commit.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionRepositorySearch performs a call to diffusion.repository.search
func (c *Conn) DiffusionRepositorySearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("diffusion.repository.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionCommitParentsQuery performs a call to diffusion.commitparentsquery
func (c *Conn) DiffusionCommitParentsQuery(req requests.CommitParentsQueryRequest) (*responses.CommitParentsQueryResponse, error) {
	var res responses.CommitParentsQueryResponse
	if err := c.Call("diffusion.commitparentsquery", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionResolveRefs performs a call to diffusion.resolverefs
func (c *Conn) DiffusionResolveRefs(req requests.ResolveRefsRequest) (*responses.ResolveRefsResponse, error) {
	var res responses.ResolveRefsResponse
	if err := c.Call("diffusion.resolverefs", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionFileContentQuery performs a call to diffusion.filecontentquery
func (c *Conn) DiffusionFileContentQuery(req requests.FileContentQueryRequest) (*responses.FileContentQueryResponse, error) {
	var res responses.FileContentQueryResponse
	if err := c.Call("diffusion.filecontentquery", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DiffusionBranchQuery performs a call to diffusion.branchquery
func (c *Conn) DiffusionBranchQuery(req requests.DiffusionBranchQueryRequest) (*responses.DiffusionBranchQueryResponse, error) {
	var res responses.DiffusionBranchQueryResponse
	if err := c.Call("diffusion.branchquery", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
