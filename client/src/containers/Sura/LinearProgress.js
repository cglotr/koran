import styled from 'styled-components'
import { LinearProgress } from '@material-ui/core'

import { colors, dimensions } from '@app/constants'

export default styled(LinearProgress)`
  &&.MuiLinearProgress-root {
    background-color: ${colors.WHITE_SMOKE};
    margin-bottom: ${dimensions.MARGIN_LARGE}px;
  }

  && .MuiLinearProgress-barColorPrimary {
    background-color: ${colors.WHITE};
  }
`
