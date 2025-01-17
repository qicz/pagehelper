/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pagehelper

const (
	DriverDummy     = "default"
	DriverMysql     = "mysql"
	DriverPostgre   = "postgre"
	DriverOracle    = "oracle"
	DriverSqlServer = "sqlserver"
)

type Modifier struct {
	OrderBy func(sql string, p *OrderByInfo) string
	Page    func(sql string, p *PageInfo) string
	Count   func(sql, countColumn string) string
}

var DummyModifier = Modifier{
	OrderBy: DummyModifyOrderBy,
	Page:    DummyModifyPage,
	Count:   DummyModifyCount,
}

func DummyModifyOrderBy(sql string, p *OrderByInfo) string {
	return sql
}

func DummyModifyPage(sql string, p *PageInfo) string {
	return sql
}

func DummyModifyCount(sql, countColumn string) string {
	return sql
}
