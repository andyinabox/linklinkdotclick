import { getLink, updateLink } from '../../lib/api'
import { Component } from '../component'

import './link.css'
export class Link extends Component {
  constructor() {
    super()
    const id = this.getAttribute('data-id')
    if (id) this.fetchData(id)
  }
  set data(d) {
    this._data = d
    this.render()
  }
  async fetchData(id) {
    this.classList.add('loading')
    this.data = await getLink(id)
    this.classList.remove('loading')
  }

  render() {
    const data = this._data

    this.slots.delete.setAttribute('data-e-id', data.id)
    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName
    this.slots.count.textContent = data.unreadCount

    this.setAttribute('data-id', data.id)
  }

  async updateLastClicked() {
    const link = { ...this._data }
    link.lastClicked = new Date().toJSON()
    console.log('update last clicked', link)
    const updatedLink = await updateLink(link)
    console.log('updated', updatedLink)
    this.data = updatedLink
  }
  connectedCallback() {
    this.onClick = () => this.updateLastClicked()
    this.addEventListener('click', this.onClick)
  }
  disconnectedCallback() {
    this.removeEventListener('click', this.onClick)
  }
}

Link.create = function (parentEl, data) {
  const linkTmpl = document.getElementById('tmpl-link')
  const linkEl = linkTmpl.content.firstElementChild.cloneNode(true)
  parentEl.appendChild(linkEl)
  linkEl.data = data
  return linkEl
}

customElements.define('linky-link', Link)
