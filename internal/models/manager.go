package models

import "github.com/aaronize/essai/pkg/db"

type ManagerOption interface {
    apply()
}

type Manager struct {
    query queryCondition
}

func NewManager() *Manager {
    return &Manager{}
}

func (m *Manager) Update(ins ...Model) (int, error) {
    if len(ins) == 0 {
        return 0, nil
    }

    tx := db.WDB().Begin()
    for _, i := range ins {
        if err := i.UpdateWithTransaction(tx); err != nil {
            tx.Rollback()
            return 0, err
        }
    }

    return len(ins), tx.Commit().Error
}

func (m *Manager) Delete(ins ...Model) (int, error) {
    if len(ins) == 0 {
        return 0, nil
    }



    return 0, nil
}

//
type queryCondition struct {
    Page    int
    Limit   string
    Order   string
    Search  string      // 模糊查询
    ConditionMap map[string]interface{} // 字段查询
}

func (qc *queryCondition) apply() {

}
