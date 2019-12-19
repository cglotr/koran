import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'
import { Typography as T } from '@material-ui/core'
import Column from './Column'

const Typography = styled(T)`
  &.MuiTypography-h3 {
    font-family: 'Scheherazade', 'Roboto', 'Helvetica', 'Arial', sans-serif;
  }
`

export default class Verse extends React.Component {
  static propTypes = {
    ayah: PropTypes.string,
    translation: PropTypes.string,
    verseNumber: PropTypes.string
  }

  render () {
    return (
      <Column>
        <Typography gutterBottom>{this.props.verseNumber}</Typography>
        <Typography align='right' variant='h3' gutterBottom>{this.props.ayah}</Typography>
        <Typography align='right' gutterBottom>{this.props.translation}</Typography>
      </Column>
    )
  }
}
