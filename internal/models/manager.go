package models

import (
    "github.com/aaronize/essai/pkg/db"
)

type Options interface {
    apply(opt *managerOpt)
}

type managerOpt struct {
    queryCondition
}

type Manager struct {
    opt     *managerOpt
}

func NewManager() *Manager {
    return &Manager{}
}

func (m *Manager) Add(ins ...Model) (int, error) {
    if len(ins) == 0 {
        return 0, nil
    }

    tx := db.WDB().Begin()
    for _, i := range ins {
        if err := i.AddWithTransaction(tx); err != nil {
            tx.Rollback()
            return 0, err
        }
    }
    return len(ins), tx.Commit().Error
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

func (m *Manager) ListAll(ins Model) ([]Model, error) {

    return nil, nil
}

func (m *Manager) GetBy(ins Model, by map[string]interface{}) ([]Model, error) {

    return nil, nil
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
