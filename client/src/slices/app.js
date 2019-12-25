import produce from 'immer'
import _ from 'lodash'
import { createSlice } from '@reduxjs/toolkit'
import { dimensions } from '@app/constants'

export default createSlice({
  initialState: {
    isDrawerOpen: false,
    isMobile: false,
    isUserDialogOpen: false,
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
    setIsUserDialogOpen: (state, action) => produce(state, (draft) => {
      const isUserDialogOpen = action.payload
      _.merge(draft, {
        isUserDialogOpen
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
