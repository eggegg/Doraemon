package bindings

type AdRequest struct {
	F string `query:"f"`
	Uid string `query:"uid"`
}

func (ar AdRequest) Validate() error {
	errs := new(RequestErrors)
	if lr.Username == "" {
		errs.Append(ErrUsernameEmpty)
	}
	if lr.Password == "" {
		errs.Append(ErrPasswordEmpty)
	}
	if errs.Len() == 0 {
		return nil
	}
	return errs
}