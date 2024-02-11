export const slotsMixin = {
  registerSlotsMixin() {
    this.slots = {}
    this.querySelectorAll('[slot]').forEach((el) => {
      this.slots[el.getAttribute('slot')] = el
    })
  },
}

export const listenMixin = {
  registerListenMixin() {
    this.listeners = []
  },
  listen(el, eventName, callback) {
    // bind the callback to this element
    const cb = callback.bind(this)
    // wrap the callback (helps for async funcs)
    const fn = (e) => cb(e)
    // add the listener
    el.addEventListener(eventName, fn)
    // add to our saved listeners
    this.listeners.push({ el, eventName, fn })
  },
  unlisten() {
    // unbind all listeners
    this.listeners.forEach(({ el, eventName, fn }) => {
      el.removeEventListener(eventName, fn)
    })
  },
}
