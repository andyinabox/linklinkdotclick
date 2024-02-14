import { renderDataMixin, eventsMixin, slotsMixin } from '../lib/mixins'
class Link extends HTMLLIElement {
  constructor() {
    super()
    this.registerSlots()
    const href = this.slots.link.getAttribute('data-href')
    if (href) {
      this.siteUrl = href
      this.removeAttribute('data-href')
    }
  }
  get data() {
    return { ...this._data }
  }
  set data(d) {
    this._data = d
    this.render()
  }
  set linkId(v) {
    this.setAttribute('data-id', v)
  }
  get linkId() {
    return parseInt(this.getAttribute('data-id'))
  }
  set siteUrl(v) {
    this.slots.link.href = v
  }
  get siteUrl() {
    return this.slots.link.href
  }
  set siteName(v) {
    this.slots.link.innerText = v
  }
  get siteName() {
    return this.slots.link.innerText
  }
  set unreadCount(v) {
    this.slots.count.innerText = v
  }
  get unreadCount() {
    return parseInt(this.slots.count.innerText)
  }
  render() {
    const { id, siteName, unreadCount, siteUrl } = this.data
    this.linkId = id
    this.siteName = siteName
    this.unreadCount = unreadCount
    this.siteUrl = siteUrl
  }
  async click(evt) {
    const now = new Date().toJSON()
    console.log('click', now)
    this.slots.form.lastClicked = now
    this.slots.form.save()
  }

  handleUpdate(evt) {
    const link = evt.detail.link
    console.log('handleUpdate', link)
    this.data = link
  }
  handleDelete(evt) {
    const id = evt.detail.id
    console.log('handleDelete', evt)
    if (this.linkId == id) {
      this.remove()
    } else {
      console.error('delete link id mismatch', id, this.linkId)
    }
  }
  connectedCallback() {
    this.listen(this.slots.link, 'click', this.click)
    this.listen(this.slots.form, 'update-link-success', this.handleUpdate)
    this.listen(this.slots.form, 'delete-link-success', this.handleDelete)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(Link.prototype, eventsMixin, slotsMixin)
customElements.define('link-list-item', Link, { extends: 'li' })
