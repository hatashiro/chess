const query = (...args) => document.querySelector(...args);

query.all = (...args) => document.querySelectorAll(...args);

export const $ = query;
