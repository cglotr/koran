import { configureStore } from '@reduxjs/toolkit'
import { quran } from '@app/slices'

export default () => {
  return configureStore({
    reducer: {
      [quran.name]: quran.reducer
    }
  })
}
