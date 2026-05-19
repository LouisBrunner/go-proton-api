package proton

import (
	"context"

	"github.com/go-resty/resty/v2"
)

func (c *Client) GetRevisionAllBlocks(ctx context.Context, shareID, linkID, revisionID string) (Revision, error) {
	var res struct {
		Revision Revision
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetResult(&res).
			Get("/drive/shares/" + shareID + "/files/" + linkID + "/revisions/" + revisionID)
	}); err != nil {
		return Revision{}, err
	}

	return res.Revision, nil
}

type RevisionVerification struct {
	VerificationCode string
	ContentKeyPacket string
}

func (c *Client) VerifyRevision(ctx context.Context, shareID, linkID, revisionID string) (RevisionVerification, error) {
	var res RevisionVerification

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.
			SetResult(&res).
			Get("/drive/shares/" + shareID + "/links/" + linkID + "/revisions/" + revisionID + "/verification")
	}); err != nil {
		return RevisionVerification{}, err
	}

	return res, nil
}

type CreatedRevision struct {
	ID string // Encrypted Revision ID
}

func (c *Client) CreateRevision(ctx context.Context, shareID, linkID string) (CreatedRevision, error) {
	var res struct {
		Revision CreatedRevision
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Post("/drive/shares/" + shareID + "/files/" + linkID + "/revisions")
	}); err != nil {
		return CreatedRevision{}, err
	}

	return res.Revision, nil
}

func (c *Client) DeleteRevision(ctx context.Context, shareID, linkID, revisionID string) error {
	return c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Delete("/drive/shares/" + shareID + "/files/" + linkID + "/revisions/" + revisionID)
	})
}

type CheckAvailableHashesReq struct {
	Hashes []string
}

type PendingHashData struct {
	Hash       []string
	RevisionID []string
	LinkID     []string
}

type AvailableHashes struct {
	AvailableHashes   []string
	PendingHashesData []PendingHashData
}

func (c *Client) CheckAvailableHashes(ctx context.Context, shareID, linkID string, req CheckAvailableHashesReq) (AvailableHashes, error) {
	var res AvailableHashes

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).SetBody(req).Post("/drive/shares/" + shareID + "/links/" + linkID + "/checkAvailableHashes")
	}); err != nil {
		return AvailableHashes{}, err
	}

	return res, nil
}
