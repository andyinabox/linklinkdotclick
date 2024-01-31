import { getLink } from '../../lib/api'
import { Component } from '../component'

import './link.css'
export class Link extends Component {
  constructor() {
    super()
    const id = this.getAttribute('data-id')
    if (id) this.fetchData(id)
  }
  async fetchData(id) {
    const data = await getLink(id)
    this.render(data)
  }
  hydrate(data) {
    this.render(data)
  }
  render(data) {
    console.log('render', data)
    this.slots.delete.setAttribute('data-e-id', data.id)
    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName
    this.slots.count.textContent = data.unreadCount

    this.setAttribute('data-id', data.id)
  }
  connectedCallback() {}
}

customElements.define('linky-link', Link)
