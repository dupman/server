/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, February 2022
 */

package sqltype

import (
	"context"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WebsiteToken string

func (t *WebsiteToken) Decrypt(privateKey string) (decrypted string, err error) {
	encryptor := helper.NewRSAEncryptor()

	err = encryptor.SetPrivateKey(privateKey)
	if err != nil {
		return decrypted, err
	}

	return encryptor.Decrypt(string(*t))
}

func (t WebsiteToken) GormValue(ctx context.Context, tx *gorm.DB) (expr clause.Expr) {
	if encryptionKey, ok := ctx.Value(constant.EncryptionKeyKey).(string); ok {
		encryptor := helper.NewRSAEncryptor()
		if err := encryptor.SetPublicKey(encryptionKey); err == nil {
			if encrypted, err := encryptor.Encrypt(string(t)); err == nil {
				return clause.Expr{SQL: "?", Vars: []interface{}{encrypted}}
			}
		}
	}

	return
}