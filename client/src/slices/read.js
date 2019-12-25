import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'

export default createSlice({
  initialState: {},
  name: 'read',
  reducers: {
    reset: () => {
      return {}
    },
    setSuraVerseRead: (state, action) => produce(state, (draft) => {
      const { isRead, suraNumber, verseNumber } = action.payload
      _.merge(draft, {
        [suraNumber]: {
          [verseNumber]: isRead
        }
      })
    })
  }
})
