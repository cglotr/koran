import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'

export default createSlice({
  initialState: {
    email: ''
  },
  name: 'user',
  reducers: {
    setUserEmail: (state, action) => produce(state, (draft) => {
      const email = action.payload
      _.merge(draft, {
        email
      })
    })
  }
})
