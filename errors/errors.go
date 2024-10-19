package errors

import "github.com/pkg/errors"

var (
	InvalidParameterError        = errors.New("invalid parameter")
	InvalidParameterValueError   = errors.New("invalid parameter value")
	UnmatchedSizeOfVectorsError  = errors.New("unmatched size of vectors")
	InvalidSizeOfMatrixError     = errors.New("invalid size of matrix")
	UnmatchedSizeOfMatricesError = errors.New("unmatched size of matrices")
	ZeroDeterminantError         = errors.New("zero determinant")
)

var (
	New          = errors.New
	Errorf       = errors.Errorf
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	WithStack    = errors.WithStack
)
