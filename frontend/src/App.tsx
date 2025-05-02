import { 
  ChakraProvider, 
  defineStyle,
  defineStyleConfig,
  extendTheme 
} from '@chakra-ui/react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { PoolProvider } from './contexts/PoolContext';
import HomePage from './pages/HomePage';
import VotePage from './pages/VotePage';

// Define theme components
const components = {
  FormLabel: defineStyleConfig({
    baseStyle: defineStyle({
      fontSize: "md",
      fontWeight: "medium",
      mb: "2px",
    }),
  }),
};

// Extend the theme with custom colors if needed
const theme = extendTheme({
  config: {
    initialColorMode: 'light',
    useSystemColorMode: false,
  },
  colors: {
    brand: {
      50: '#e0f7fa',
      100: '#b2ebf2',
      200: '#80deea',
      300: '#4dd0e1',
      400: '#26c6da',
      500: '#00bcd4',
      600: '#00acc1',
      700: '#0097a7',
      800: '#00838f',
      900: '#006064',
    },
  },
  components,
});

function App() {
  return (
    <ChakraProvider theme={theme}>
      <PoolProvider>
        <Router>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/vote/:poolId" element={<VotePage />} />
          </Routes>
        </Router>
      </PoolProvider>
    </ChakraProvider>
  );
}

export default App;