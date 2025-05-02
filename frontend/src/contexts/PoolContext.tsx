import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { Pool, getPool, voteOnPool, initializePoolsFromConfig } from '../services/api';
import config from '../config';

interface PoolContextType {
  currentPool: Pool | null;
  loading: boolean;
  error: string | null;
  voted: boolean;
  refreshPool: () => Promise<void>;
  vote: (optionId: string) => Promise<void>;
  nextPool: () => void;
  previousPool: () => void;
}

const PoolContext = createContext<PoolContextType | undefined>(undefined);

interface PoolProviderProps {
  children: ReactNode;
}

export const PoolProvider: React.FC<PoolProviderProps> = ({ children }) => {
  const [currentPool, setCurrentPool] = useState<Pool | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [voted, setVoted] = useState<boolean>(false);
  const [currentIndex, setCurrentIndex] = useState<number>(config.currentPoolIndex);

  // Initialize pools from config
  useEffect(() => {
    const initializePools = async () => {
      try {
        await initializePoolsFromConfig();
        await loadCurrentPool();
      } catch (e) {
        setError('Failed to initialize pools');
        console.error('Error initializing pools:', e);
        setLoading(false);
      }
    };

    initializePools();
  }, []);

  // Set up refresh interval
  useEffect(() => {
    if (!currentPool) return;

    const intervalId = setInterval(() => {
      refreshPool();
    }, config.refreshInterval * 1000);

    return () => clearInterval(intervalId);
  }, [currentPool]);

  const loadCurrentPool = async () => {
    setLoading(true);
    setError(null);
    setVoted(false);

    try {
      const poolConfig = config.pools[currentIndex];
      if (!poolConfig || !poolConfig.id) {
        setError('No active pool found');
        setCurrentPool(null);
        setLoading(false);
        return;
      }

      const pool = await getPool(poolConfig.id);
      setCurrentPool(pool);
      
      // Check if user has voted in local storage
      const votedPools = JSON.parse(localStorage.getItem('votedPools') || '{}');
      setVoted(!!votedPools[pool.id]);
    } catch (e) {
      setError('Failed to load pool');
      console.error('Error loading pool:', e);
    } finally {
      setLoading(false);
    }
  };

  const refreshPool = async () => {
    if (!currentPool) return;
    
    try {
      const updatedPool = await getPool(currentPool.id);
      setCurrentPool(updatedPool);
    } catch (e) {
      console.error('Error refreshing pool:', e);
    }
  };

  const vote = async (optionId: string) => {
    if (!currentPool) return;
    
    setLoading(true);
    setError(null);
    
    try {
      // Vote with a random user ID
      const updatedPool = await voteOnPool(currentPool.id, optionId);
      setCurrentPool(updatedPool);
      setVoted(true);
      
      // Save vote in local storage to prevent revoting
      const votedPools = JSON.parse(localStorage.getItem('votedPools') || '{}');
      votedPools[currentPool.id] = true;
      localStorage.setItem('votedPools', JSON.stringify(votedPools));
    } catch (e) {
      setError('Failed to submit vote');
      console.error('Error voting:', e);
    } finally {
      setLoading(false);
    }
  };

  const nextPool = () => {
    if (currentIndex < config.pools.length - 1) {
      setCurrentIndex(currentIndex + 1);
      loadCurrentPool();
    }
  };

  const previousPool = () => {
    if (currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
      loadCurrentPool();
    }
  };

  return (
    <PoolContext.Provider
      value={{
        currentPool,
        loading,
        error,
        voted,
        refreshPool,
        vote,
        nextPool,
        previousPool
      }}
    >
      {children}
    </PoolContext.Provider>
  );
};

export const usePool = () => {
  const context = useContext(PoolContext);
  if (context === undefined) {
    throw new Error('usePool must be used within a PoolProvider');
  }
  return context;
};

export default PoolContext;