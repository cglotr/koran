import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'

export default createSlice({
  initialState: {
    id: 0,
    email: '',
    uid: '',
    token: ''
  },
  name: 'user',
  reducers: {
    requestSignOut: (state) => state,
    reset: (state) => produce(state, (draft) => {
      _.merge(draft, {
        id: 0,
        uid: '',
        token: ''
      })
    }),
    setUser: (state, action) => produce(state, (draft) => {
      const {
        id,
        uid,
        token
      } = action.payload
      _.merge(draft, {
        id,
        uid,
        token
      })
    }),
    setUserEmail: (state, action) => produce(state, (draft) => {
      const email = action.payload
      _.merge(draft, {
        email
      })
    })
  }
})
