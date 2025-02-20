// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"context"
	"fmt"

	"github.com/harness/gitness/app/api/usererror"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

// Find finds a webhook from the provided repository.
func (c *Controller) Find(
	ctx context.Context,
	session *auth.Session,
	repoRef string,
	webhookID int64,
) (*types.Webhook, error) {
	repo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoView)
	if err != nil {
		return nil, err
	}

	return c.getWebhookVerifyOwnership(ctx, repo.ID, webhookID)
}

func (c *Controller) getWebhookVerifyOwnership(ctx context.Context, repoID int64,
	webhookID int64) (*types.Webhook, error) {
	if webhookID <= 0 {
		return nil, usererror.BadRequest("A valid webhook ID must be provided.")
	}

	webhook, err := c.webhookStore.Find(ctx, webhookID)
	if err != nil {
		return nil, fmt.Errorf("failed to find webhook with id %d: %w", webhookID, err)
	}

	// ensure the webhook actually belongs to the repo
	if webhook.ParentType != enum.WebhookParentRepo || webhook.ParentID != repoID {
		return nil, fmt.Errorf("webhook doesn't belong to requested repo. Returning error %w", usererror.ErrNotFound)
	}

	return webhook, nil
}
