import axios from 'axios';
import config from '../config';
import { v4 as uuidv4 } from 'uuid';

const api = axios.create({
  baseURL: config.apiUrl,
  headers: {
    'Content-Type': 'application/json',
  },
});

export interface Pool {
  id: string;
  title: string;
  description: string;
  options: Option[];
  created_at: string;
  expires_at: string;
}

export interface Option {
  id: string;
  text: string;
  votes: number;
}

export interface Vote {
  option_id: string;
}

// Get a specific pool
export const getPool = async (id: string): Promise<Pool> => {
  const response = await api.get(`/pools/${id}`);
  return response.data;
};

// Get all pools
export const getPools = async (): Promise<Pool[]> => {
  const response = await api.get('/pools');
  return response.data;
};

// Vote on a pool with a random user ID
export const voteOnPool = async (
  poolId: string,
  optionId: string
): Promise<Pool> => {
  // Generate a random UUID for the vote
  const voteId = uuidv4();
  
  const response = await api.post(`/pools/${poolId}/vote`, {
    user_id: voteId,
    option_id: optionId,
  });
  return response.data;
};

// Initialize pools from config if they don't exist
export const initializePoolsFromConfig = async (): Promise<void> => {
  try {
    // Get all existing pools
    const existingPools = await getPools();
    
    // Create pools from config if they don't already exist
    for (let i = 0; i < config.pools.length; i++) {
      const configPool = config.pools[i];
      
      // Check if this pool already exists by title (simple matching)
      const existingPool = existingPools.find(
        (p) => p.title.toLowerCase() === configPool.title.toLowerCase()
      );
      
      if (existingPool) {
        // Store the ID in our config for future reference
        config.pools[i].id = existingPool.id;
      } else if (configPool.active) {
        // Create the pool if it doesn't exist and is set as active
        try {
          const newPool = await createPool(
            configPool.title,
            configPool.description,
            configPool.options,
            configPool.expiresIn
          );
          
          // Store the new ID in our config
          config.pools[i].id = newPool.id;
        } catch (err) {
          console.error(`Failed to create pool: ${configPool.title}`, err);
        }
      }
    }
  } catch (err) {
    console.error('Failed to initialize pools from config', err);
  }
};

// Create a new pool
export const createPool = async (
  title: string,
  description: string,
  options: string[],
  expiresIn: number
): Promise<Pool> => {
  const response = await api.post('/pools', {
    title,
    description,
    options,
    expires_in: expiresIn,
  });
  return response.data;
};

export default api;