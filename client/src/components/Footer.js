import React from 'react'
import styled from 'styled-components'
import { Typography } from '@material-ui/core'

import { Column } from '@app/components'
import { dimensions } from '@app/constants'

const MyColumn = styled(Column)`
  margin-bottom: ${dimensions.LENGTH_100}px;
  margin-top: ${dimensions.LENGTH_100}px;
`

export default class Footer extends React.Component {
  render () {
    return (
      <MyColumn>
        <Typography align='center' variant='h5'>☪️</Typography>
      </MyColumn>
    )
  }
}
