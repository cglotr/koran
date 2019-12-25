import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'
import { Checkbox, Typography as T } from '@material-ui/core'
import Column from './Column'
import Row from './Row'

const Typography = styled(T)`
  &.MuiTypography-h3 {
    font-family: 'Scheherazade', 'Roboto', 'Helvetica', 'Arial', sans-serif;
  }
`

export default class Verse extends React.Component {
  static propTypes = {
    ayah: PropTypes.string,
    isRead: PropTypes.bool,
    onCheckboxChange: PropTypes.func.isRequired,
    translation: PropTypes.string,
    verseNumber: PropTypes.string
  }

  render () {
    const translation = this.getTranslation(this.props.translation)
    return (
      <Column>
        <Row justifyContent='flex-end'>
          <Typography>{this.props.verseNumber}</Typography>
          <Checkbox
            checked={this.props.isRead}
            color='default'
            onChange={(e) => this.props.onCheckboxChange(e.target.checked)}
            value='primary'
          />
        </Row>
        <Row justifyContent='flex-end'>
          <Typography align='right' gutterBottom variant='h3'>{this.props.ayah}</Typography>
        </Row>
        <Row>
          <Typography align='justify' gutterBottom>{translation}</Typography>
        </Row>
      </Column>
    )
  }

  getTranslation = (translation) => {
    if (!translation) return ''
    return translation.replace(/&quot;/g, '"')
  }

  handleCheckboxChange = (isRead) => {
    this.props.onCheckboxChange(isRead)
  }
}
