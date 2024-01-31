import { getLink } from '../../lib/api'
import { Component } from '../component'

import './link.css'
export class Link extends Component {
  constructor() {
    super()
  }
  // async getData() {
  //   const id = this.getAttribute('data-id')
  //   this.data = await getLink(id)
  //   this.render()
  // }
  render() {}
  connectedCallback() {
    // this.getData()
  }
}

customElements.define('linky-link', Link)
