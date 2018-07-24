package bindings

type AdRequest struct {
	F string `query:"f"`
	Uid string `query:"uid"`
}

func (ar AdRequest) Validate() error {
	errs := new(RequestErrors)
	if ar.F == "" {
		errs.Append(ErrFEmpty)
	}
	if ar.Uid == "" {
		errs.Append(ErrUidEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}