export interface PoolConfig {
  id?: string;        // Pool ID (from backend) - empty if not yet created
  title: string;      // Pool title
  description: string; // Pool description
  options: string[];  // Available options
  expiresIn: number;  // Hours until expiration
  active: boolean;    // If the pool is currently active
}

export interface AppConfig {
  apiUrl: string;           // Backend API URL
  pools: PoolConfig[];      // List of pools to display
  currentPoolIndex: number; // Index of currently active pool
  showResults: boolean;     // Whether to show results immediately
  refreshInterval: number;  // Refresh interval in seconds
}

const config: AppConfig = {
  apiUrl: import.meta.env.VITE_API_URL || 'http://localhost:8081/api',
  pools: [
    {
      title: "Favorite Programming Language",
      description: "What is your favorite programming language?",
      options: [
        "JavaScript/TypeScript", 
        "Python", 
        "Go", 
        "Java", 
        "C#", 
        "Rust"
      ],
      expiresIn: 24,
      active: true
    },
    {
      title: "Best Frontend Framework",
      description: "Which frontend framework do you prefer?",
      options: [
        "React", 
        "Vue", 
        "Angular", 
        "Svelte", 
        "Solid"
      ],
      expiresIn: 24,
      active: false
    },
    {
      title: "Preferred Backend Language",
      description: "What's your go-to backend language?",
      options: [
        "Node.js", 
        "Python", 
        "Go", 
        "Java", 
        "Ruby", 
        "PHP", 
        "C#/.NET"
      ],
      expiresIn: 24,
      active: false
    }
  ],
  currentPoolIndex: 0,
  showResults: true,
  refreshInterval: 5
};

export default config;
