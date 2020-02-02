import { Scrollbars as S } from 'react-custom-scrollbars'
import styled from 'styled-components'
import { Container as C } from '@material-ui/core'

import { dimensions } from '@app/constants'

const Container = styled(C)`
height: 100%;
`

const Page = styled.div`
display: flex;
flex: 1;
padding-left: ${(props) => props.paddingLeft}px;
padding-top: ${dimensions.APP_BAR_HEIGHT}px;
`

const Scrollbars = styled(S)`
flex: 1;
`

export {
  Container,
  Page,
  Scrollbars
}
