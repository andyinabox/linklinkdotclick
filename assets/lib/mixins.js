export const slotsMixin = {
  registerSlots() {
    this.slots = {}
    this.querySelectorAll('[slot]').forEach((el) => {
      this.slots[el.getAttribute('slot')] = el
    })
  },
}

export const eventsMixin = {
  listen(el, eventName, callback) {
    if (!this.listeners) this.listeners = []

    // bind the callback to this element
    const cb = callback.bind(this)
    // wrap the callback (helps for async funcs)
    const fn = (e) => cb(e)
    // add the listener
    el.addEventListener(eventName, fn)
    // add to our saved listeners
    this.listeners.push({ el, eventName, fn })
  },
  unlistenAll() {
    if (!this.listeners) return

    // unbind all listeners
    this.listeners.forEach(({ el, eventName, fn }) => {
      el.removeEventListener(eventName, fn)
    })
  },
  broadcast(name, detail = {}) {
    this.dispatchEvent(
      new CustomEvent(name, {
        detail,
        bubbles: true,
      })
    )
  },
}
