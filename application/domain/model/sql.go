package model

type SQL struct {
	Query string
}

func (s *SQL) GetSQL(ts []*Schema) *SQL {
	sql := "create table aaaa ("

	i := len(ts) - 1
	for s, t := range ts {
		if s != 0 {
			sql += t.LogicalName + " "
			sql += t.Type

			if s < i {
				sql += ", "
			}
		}
	}
	sql += ");"
	return s.setQuery(sql)
}

func (s *SQL) setQuery(st string) *SQL {
	s.Query = st
	return s
}
