package storage

import (
	"database/sql"
	"fmt"

	"github.com/chesnovsky/fliboobstier/bot_helpers"
	"github.com/chesnovsky/fliboobstier/logger"
	_ "github.com/mattn/go-sqlite3"
)

// Database-dependent function
// Write a new one if you are moving to another database
func NewSqLite(url string) (StorageInstance, error) {
	storageInstance := StorageInstance{url, nil}
	err := error(nil)
	storageInstance.db, err = sql.Open("sqlite3", url)
	return storageInstance, err
}

// Database-dependent function
// Write a new one if you are moving to another database
func (storage StorageInstance) GetRegexActionElementsSqLite(action_id string, element_type string) ([]string, error) {
	Elements := []string{}
	err_list := []string{}
	query_str := fmt.Sprintf("select element_id from regex_action_%ss where action_id=?", element_type)
	rows, err := storage.db.Query(query_str, action_id)
	if err != nil {
		err_msg := fmt.Sprintf("Cannot select content <%s> for action <%s> from database, query <%s>: %+v", element_type, action_id, query_str, err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
		return Elements, bot_helpers.ErrListToError(err_list)
	}
	defer rows.Close()
	for rows.Next() {
		var content_id string
		err = rows.Scan(&content_id)
		if err != nil {
			err_msg := fmt.Sprintf("Cannot read selected row into data struct: %+v", err)
			logger.Logger.Error(err_msg)
			err_list = append(err_list, err_msg)
		} else {
			Elements = append(Elements, content_id)
		}
	}
	if rows.Err() != nil {
		err_msg := fmt.Sprintf("WTF while reading rows from database: %+v", rows.Err())
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}
	return Elements, bot_helpers.ErrListToError(err_list)
}

// Database-dependent function
// Write a new one if you are moving to another database
func (storage StorageInstance) AddRegexActionElementSqlite(action_id string, element_type string, element_id string) error {
	query_str := fmt.Sprintf("insert into regex_action_%ss(action_id, element_id) values (?, ?)", element_type)
	result, err := storage.db.Exec(query_str, action_id, element_id)
	if err != nil {
		return err
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error while reading affected_rows for query <%s>, params <%s>, <%s>: %+v", query_str, action_id, element_id, err)
	}
	if rows_affected != 1 {
		return fmt.Errorf("Query <%s> hasn't changed anything, params: <%s>, <%s>", query_str, action_id, element_id)
	}
	return nil
}

// Database-dependent function
// Write a new one if you are moving to another database
func (storage StorageInstance) RemoveRegexActionElementSqlite(action_id string, element_type string, element_id string) error {
	query_str := fmt.Sprintf("delete from regex_action_%ss where action_id=? and element_id=?", element_type)
	result, err := storage.db.Exec(query_str, action_id, element_id)
	if err != nil {
		return err
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error while reading affected_rows for query <%s>, params <%s>, <%s>: %+v", query_str, action_id, element_id, err)
	}
	if rows_affected != 1 {
		return fmt.Errorf("Query <%s> hasn't changed anything, params: <%s>, <%s>", query_str, action_id, element_id)
	}
	return nil
}

// Database-dependent function
// Write a new one if you are moving to another database
func (storage StorageInstance) GetAdminListSqlite() ([]string, error) {
	AdminList := []string{}
	err_list := []string{}
	query_str := fmt.Sprintf("select admin_login from admins")
	rows, err := storage.db.Query(query_str)
	if err != nil {
		err_msg := fmt.Sprintf("Cannot select admin list from database, query <%s>: %+v", query_str, err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
		return AdminList, bot_helpers.ErrListToError(err_list)
	}
	defer rows.Close()
	for rows.Next() {
		var admin_login string
		err = rows.Scan(&admin_login)
		if err != nil {
			err_msg := fmt.Sprintf("Cannot read selected row into data struct: %+v", err)
			logger.Logger.Error(err_msg)
			err_list = append(err_list, err_msg)
		} else {
			AdminList = append(AdminList, admin_login)
		}
	}
	if rows.Err() != nil {
		err_msg := fmt.Sprintf("WTF while reading rows from database: %+v", rows.Err())
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}
	return AdminList, bot_helpers.ErrListToError(err_list)
}
