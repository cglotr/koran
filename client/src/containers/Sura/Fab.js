import styled from 'styled-components'
import { Fab } from '@material-ui/core'

import { dimensions } from '@app/constants'

export default styled(Fab)`
  &.MuiFab-root {
    bottom: 0;
    margin-bottom: ${dimensions.MARGIN_XLARGE}px;
    margin-right: ${dimensions.MARGIN_XLARGE}px;
    position: fixed;
    right: 0;
  }
`
