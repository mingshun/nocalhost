/*
Copyright 2020 The Nocalhost Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster_user

import (
	"context"
	"nocalhost/internal/nocalhost-api/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ClusterUserRepo interface {
	Create(ctx context.Context, model model.ClusterUserModel) (uint64, error)
	Close()
}

type clusterUserRepo struct {
	db *gorm.DB
}

func NewApplicationClusterRepo(db *gorm.DB) ClusterUserRepo {
	return &clusterUserRepo{
		db: db,
	}
}

func (repo *clusterUserRepo) Create(ctx context.Context, model model.ClusterUserModel) (id uint64, err error) {
	err = repo.db.Create(&model).Error
	if err != nil {
		return 0, errors.Wrap(err, "[application_cluster_repo] create application_cluster error")
	}

	return model.ID, nil
}

// Close close db
func (repo *clusterUserRepo) Close() {
	repo.db.Close()
}
