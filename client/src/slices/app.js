import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'
import { dimensions } from '@app/constants'

export default createSlice({
  initialState: {
    isDrawerOpen: false,
    isMobile: false,
    width: 0
  },
  name: 'app',
  reducers: {
    requestSignIn: (state) => state,
    setIsDrawerOpen: (state, action) => produce(state, (draft) => {
      const isDrawerOpen = action.payload
      _.merge(draft, {
        isDrawerOpen
      })
    }),
    setWidth: (state, action) => produce(state, (draft) => {
      const width = action.payload
      _.merge(draft, {
        isMobile: width <= dimensions.BREAKPOINT_MOBILE,
        width
      })
    })
  }
})
