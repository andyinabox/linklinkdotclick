class APIError extends Error {
  constructor(status, data) {
    const { error } = data
    let err = error
    if (!err) {
      err = new Error('unknown API error')
    }
    super(`${status}: ${err.toString()}`)
    this.status = status
    this.responseBody = data
  }
}

export const apiCall = async function (endpoint, method, body) {
  const options = {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
    mode: 'cors',
  }
  if (method !== 'GET') {
    options.body = JSON.stringify(body)
  }
  const response = await fetch('/api' + endpoint, options)
  const data = await response.json()
  if (data.success) {
    return data
  } else {
    throw new APIError(response.status, data)
  }
}

//
// links
//
export const getLinks = async function () {
  const { payload } = await apiCall('/links', 'GET')
  if (error) {
    throw error
  }
  return payload
}

export const createLink = async function (url) {
  const { payload } = await apiCall('/links', 'POST', { url })
  return payload
}

export const getLink = async function (id, refresh = true) {
  let endpoint = '/links/' + id
  if (refresh) endpoint += '?refresh'
  const { payload } = await apiCall(endpoint, 'GET')
  return payload
}
export const updateLink = async function (link, refresh = true) {
  let endpoint = '/links/' + link.id
  if (refresh) endpoint += '?refresh'
  const { payload } = await apiCall(endpoint, 'PUT', link)
  return payload
}

export const deleteLink = async function (id) {
  const { payload } = await apiCall('/links/' + id, 'DELETE')
  return payload
}

//
// users
//
export const updateSelf = async function (user) {
  let endpoint = '/self'
  const { payload } = await apiCall(endpoint, 'PUT', user)
  return payload
}
