import { eventsMixin } from '../lib/mixins'
export class LinkList extends HTMLOListElement {
  constructor() {
    super()
    this.template = this.querySelector('template')
    this.items = this.querySelector('li')
  }
  handleLinkAddSuccess(evt) {
    const link = evt.detail.link
    console.log('need to add link to dom', link)
  }
  connectedCallback() {
    this.listen(document, 'link-add-success', this.handleLinkAddSuccess)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}

Object.assign(LinkList.prototype, eventsMixin)
customElements.define('link-list', LinkList, { extends: 'ol' })
