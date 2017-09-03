package main

func findWorks(works *[]Work) *Error {
	if err := db.Select(works); err != nil {
		return systemError
	}

	if len(*works) == 0 {
		return notFound
	}

	return nil
}

func findWork(w *Work, id interface{}) *Error {
	var works []Work
	if err := db.Select(&works, db.Where("id", "=", id)); err != nil {
		return systemError
	}

	if len(works) == 0 {
		return notFound
	}

	*w = works[0]

	return nil
}

func saveWork(w *Work) *Error {
	if w.ID != 0 {
		if _, err := db.Update(w); err != nil {
			return systemError
		}

		return nil
	}

	if _, err := db.Insert(w); err != nil {
		return systemError
	}

	return nil
}

func deleteWork(w *Work) *Error {
	if _, err := db.Delete(w); err != nil {
		return systemError
	}

	return nil
}
