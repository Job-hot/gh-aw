// Chaos test: split-file-refactorer
// Refactored utility functions with JSDoc

/**
 * Format a date object to ISO string
 */
function formatDate(date) {
  return date.toISOString();
}

function parseJSON(str) {
  return JSON.parse(str);
}

module.exports = { formatDate, parseJSON };
