package ratelimiter

// const (
// 	defaultLimitDur = 1 * time.Second
// 	defaultReqLimit = 100
// )

// type ClientID uint64

// // ReqStorage is an abstraction for request storage
// // it could be implemented using redis or inmemory storage
// type ReqStorage interface {
// 	Inc(clientID ClientID) error
// 	Sum(d time.Duration, clientID ClientID) (int, error)
// }

// type RateLimiter interface {
// 	IsAllow(clientID ClientID) (bool, error)
// }

// func NewRateLimiter(limitDur time.Duration, reqLimit int, rs ReqStorage) RateLimiter {
// 	if rs == nil {
// 		// fallback to some default implementation
// 	}
// 	if limitDur == 0 {
// 		limitDur = defaultLimitDur
// 	}
// 	if reqLimit == 0 {
// 		reqLimit = defautlReqLimit
// 	}

// 	return &rateLimiter{
// 		limitDur:   limitDur,
// 		rateLimit:  reqLimit,
// 		reqstorage: rs,
// 	}
// }

// type rateLimiter struct {
// 	limitDur   time.Duration
// 	rateLimit  int
// 	reqstorage ReqStorage
// }

// func (r *rateLimiter) IsAllow(clientID ClientID) (bool, error) {
// 	if clientID == 0 {
// 		return false, nil
// 	}

// 	reqNum, err := r.reqstorage.Sum(r.limitDur, clientID)
// 	if err != nil {
// 		return false, err
// 	}

// 	if reqNum > r.rateLimit {
// 		return false, nil
// 	}

// 	err = r.reqstorage.Inc(clientID)
// 	if err != nil {
// 		log.Error("err incrementing usage")
// 	}

// 	return true, nil
// }
