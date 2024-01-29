import { Component } from '../component'
import './link.css'
class Link extends Component {
  async getData() {
    const id = this.getAttribute('data-id')
    const response = await fetch(`/api/links/${id}?refresh`)
    const json = await response.json()
    this.data = json
    this.render()
  }
  render() {
    if (this.slots.count) {
      this.slots.count.innerHTML = `(${this.data.unreadCount})`
    }
  }
  connectedCallback() {
    this.getData()
  }
}
customElements.define('linky-link', Link)
