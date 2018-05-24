//
// command_generic.go
// Copyright (C) 2018 YanMing <yming0221@gmail.com>
//
// Distributed under terms of the MIT license.
//

package server

import (
	"github.com/yongman/go/util"
	"github.com/yongman/tidis/terror"
	"github.com/yongman/tidis/tidis"
)

func pexpireGeneric(c *Client, t byte) error {
	if len(c.args) != 2 {
		return terror.ErrCmdParams
	}
	var (
		v   int
		err error
	)

	i, err := util.StrBytesToInt64(c.args[1])
	if err != nil {
		return terror.ErrCmdParams
	}

	switch t {
	case tidis.TSTRING:
		if !c.IsTxn() {
			v, err = c.tdb.PExpire(c.args[0], i)
		} else {
			v, err = c.tdb.PExpireWithTxn(c.GetCurrentTxn(), c.args[0], i)
		}
	case tidis.TLISTMETA:
		v, err = c.tdb.LPExpire(c.args[0], i)
	case tidis.THASHMETA:
		v, err = c.tdb.HPExpire(c.args[0], i)
	case tidis.TSETMETA:
		v, err = c.tdb.SPExpire(c.args[0], i)
	case tidis.TZSETMETA:
		v, err = c.tdb.ZPExpire(c.args[0], i)
	}
	if err != nil {
		return err
	}

	return c.Resp(int64(v))
}

func pexpireatGeneric(c *Client, t byte) error {
	if len(c.args) != 2 {
		return terror.ErrCmdParams
	}

	var (
		v   int
		err error
	)

	i, err := util.StrBytesToInt64(c.args[1])
	if err != nil {
		return terror.ErrCmdParams
	}

	switch t {
	case tidis.TSTRING:
		if !c.IsTxn() {
			v, err = c.tdb.PExpireAt(c.args[0], i)
		} else {
			v, err = c.tdb.PExpireAtWithTxn(c.GetCurrentTxn(), c.args[0], i)
		}
	case tidis.TLISTMETA:
		v, err = c.tdb.LPExpireAt(c.args[0], i)
	case tidis.THASHMETA:
		v, err = c.tdb.HPExpireAt(c.args[0], i)
	case tidis.TSETMETA:
		v, err = c.tdb.SPExpireAt(c.args[0], i)
	case tidis.TZSETMETA:
		v, err = c.tdb.ZPExpireAt(c.args[0], i)

	}
	if err != nil {
		return err
	}

	c.Resp(int64(v))

	return nil
}

func expireGeneric(c *Client, t byte) error {
	if len(c.args) != 2 {
		return terror.ErrCmdParams
	}

	var (
		v   int
		err error
	)

	i, err := util.StrBytesToInt64(c.args[1])
	if err != nil {
		return terror.ErrCmdParams
	}

	switch t {
	case tidis.TSTRING:
		if !c.IsTxn() {
			v, err = c.tdb.Expire(c.args[0], i)
		} else {
			v, err = c.tdb.ExpireWithTxn(c.GetCurrentTxn(), c.args[0], i)
		}
	case tidis.TLISTMETA:
		v, err = c.tdb.LExpire(c.args[0], i)
	case tidis.THASHMETA:
		v, err = c.tdb.HExpire(c.args[0], i)
	case tidis.TSETMETA:
		v, err = c.tdb.SExpire(c.args[0], i)
	case tidis.TZSETMETA:
		v, err = c.tdb.ZExpire(c.args[0], i)
	}
	if err != nil {
		return err
	}

	return c.Resp(int64(v))
}

func expireatGeneric(c *Client, t byte) error {
	if len(c.args) != 2 {
		return terror.ErrCmdParams
	}
	var (
		v   int
		err error
	)
	i, err := util.StrBytesToInt64(c.args[1])
	if err != nil {
		return terror.ErrCmdParams
	}

	switch t {
	case tidis.TSTRING:
		if !c.IsTxn() {
			v, err = c.tdb.ExpireAt(c.args[0], i)
		} else {
			v, err = c.tdb.ExpireAtWithTxn(c.GetCurrentTxn(), c.args[0], i)
		}
	case tidis.TLISTMETA:
		v, err = c.tdb.LExpireAt(c.args[0], i)
	case tidis.THASHMETA:
		v, err = c.tdb.HExpireAt(c.args[0], i)
	case tidis.TSETMETA:
		v, err = c.tdb.SExpireAt(c.args[0], i)
	case tidis.TZSETMETA:
		v, err = c.tdb.ZExpireAt(c.args[0], i)
	}

	return c.Resp(int64(v))
}

func pttlGeneric(c *Client, t byte) error {
	if len(c.args) != 1 {
		return terror.ErrCmdParams
	}

	var (
		v   int64
		err error
	)

	switch t {
	case tidis.TSTRING:
		v, err = c.tdb.PTtl(c.GetCurrentTxn(), c.args[0])
	case tidis.TLISTMETA:
		v, err = c.tdb.LPTtl(c.args[0])
	case tidis.THASHMETA:
		v, err = c.tdb.HPTtl(c.args[0])
	case tidis.TSETMETA:
		v, err = c.tdb.SPTtl(c.args[0])
	case tidis.TZSETMETA:
		v, err = c.tdb.ZPTtl(c.args[0])
	}
	if err != nil {
		return err
	}

	return c.Resp(v)
}

func ttlGeneric(c *Client, t byte) error {
	if len(c.args) != 1 {
		return terror.ErrCmdParams
	}

	var (
		v   int64
		err error
	)

	switch t {
	case tidis.TSTRING:
		v, err = c.tdb.Ttl(c.GetCurrentTxn(), c.args[0])
	case tidis.TLISTMETA:
		v, err = c.tdb.LTtl(c.args[0])
	case tidis.THASHMETA:
		v, err = c.tdb.HTtl(c.args[0])
	case tidis.TSETMETA:
		v, err = c.tdb.STtl(c.args[0])
	case tidis.TZSETMETA:
		v, err = c.tdb.ZTtl(c.args[0])
	}
	if err != nil {
		return err
	}

	return c.Resp(v)
}