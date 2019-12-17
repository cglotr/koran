import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'

export default createSlice({
  initialState: {
    isDrawerOpen: false,
    isMobile: false,
    width: 0
  },
  name: 'app',
  reducers: {
    setIsDrawerOpen: (state, action) => produce(state, (draft) => {
      const isDrawerOpen = action.payload
      _.merge(draft, {
        isDrawerOpen
      })
    }),
    setWidth: (state, action) => produce(state, (draft) => {
      const width = action.payload
      _.merge(draft, {
        isMobile: width <= 666,
        width
      })
    })
  }
})
