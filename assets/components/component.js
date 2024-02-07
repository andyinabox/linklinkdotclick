export class Component extends HTMLElement {
  constructor() {
    super()
    // create a collection of elements with the
    // "slot" attribute
    this.slots = {}
    this.querySelectorAll('[slot]').forEach((el) => {
      this.slots[el.getAttribute('slot')] = el
    })
    this.listeners = []
  }
  // automatically render when data is set
  set data(d) {
    this._data = d
    this.render()
  }

  // create a clone of data when retrieving
  // (we only want this changed when using setter)
  get data() {
    return { ...this._data }
  }
  set loading(bool) {
    if (bool) {
      this.classList.add('loading')
    } else {
      this.classList.remove('loading')
    }
  }
  get loading() {
    return this.classList.contains('loading')
  }
  render() {}
  listen(el, eventName, callback) {
    // bind the callback to this element
    const cb = callback.bind(this)
    // wrap the callback (helps for async funcs)
    const fn = (e) => cb(e)
    // add the listener
    el.addEventListener(eventName, fn)
    // add to our saved listeners
    this.listeners.push({ el, eventName, fn })
  }
  unlisten() {
    // unbind all listeners
    this.listeners.forEach(({ el, eventName, fn }) => {
      el.removeEventListener(eventName, fn)
    })
  }
  disconnectedCallback() {
    this.unlisten()
  }
}
