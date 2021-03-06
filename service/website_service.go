/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package service

import (
	"context"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/helper"
	"github.com/dupman/server/model"
	"github.com/dupman/server/repository"
	"github.com/google/uuid"
)

// WebsiteService data structure.
type WebsiteService struct {
	AbstractService
	repository repository.WebsiteRepository
}

// NewWebsiteService creates a new WebsiteService.
func NewWebsiteService(repository repository.WebsiteRepository) WebsiteService {
	return WebsiteService{
		repository: repository,
	}
}

// Create creates new website.
func (s WebsiteService) Create(website *model.Website, encryptionKey string) (err error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.EncryptionKeyKey, encryptionKey)

	return s.repository.WithContext(ctx).Create(website).Error
}

// GetAll gets all websites.
func (s WebsiteService) GetAll(pagination *helper.Pagination) (websites []model.Website, err error) {
	return websites, s.repository.
		Scopes(s.paginate(websites, pagination, s.repository.DB)).
		Find(&websites).
		Error
}

// GetByUser gets all websites for given user.
func (s WebsiteService) GetByUser(
	userID uuid.UUID,
	pagination *helper.Pagination,
) (websites []model.Website, err error) {
	return websites, s.repository.
		Scopes(s.withUser(userID)).
		Scopes(s.paginate(websites, pagination, s.repository.DB)).
		Find(&websites).
		Error
}
