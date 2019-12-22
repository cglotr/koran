import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'

const D = styled.div`
  align-items: center;
  display: flex;
  flex: 1;
  flex-direction: row;
  justify-content: ${(props) => props.justifyContent};
`

export default class Row extends React.Component {
  static propTypes = {
    children: PropTypes.any.isRequired
  }

  render () {
    return (
      <D {...this.props}>{this.props.children}</D>
    )
  }
}
