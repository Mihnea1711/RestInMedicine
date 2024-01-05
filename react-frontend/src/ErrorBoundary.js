import React, { Component } from 'react';

class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    return { hasError: true };
  }

  componentDidCatch(error, errorInfo) {
    // Log the error to your error tracking service (e.g., Sentry, LogRocket)
    console.error('Error caught by boundary:', error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return (
        <div>
          <h2>Something went wrong.</h2>
          <p>Please try refreshing the page or contact support if the issue persists.</p>
        </div>
      );
    }
  
    return this.props.children;
  }
  
}

export default ErrorBoundary;
