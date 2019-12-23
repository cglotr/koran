import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'

const D = styled.div`
  align-items: stretch;
  align-items: ${(props) => props.alignItems};
  display: flex;
  flex: 1;
  flex-direction: column;
  height: 100%;
  justify-content: ${(props) => props.justifyContent};
`

export default class Column extends React.Component {
  static propTypes = {
    children: PropTypes.any
  }

  render () {
    return (
      <D {...this.props}>{this.props.children}</D>
    )
  }
}
