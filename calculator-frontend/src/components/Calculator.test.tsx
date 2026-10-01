import { render, screen, fireEvent } from '@testing-library/react';
import { describe, beforeEach, expect, test, vi } from 'vitest'
import '@testing-library/jest-dom';
import Calculator from './Calculator';

describe('Calculator Component', () => {
  beforeEach(() => {
    render(<Calculator />);
  });

  test('renders calculator component', () => {
    const calculatorElement = screen.getByTestId("calculator");
    expect(calculatorElement).toBeInTheDocument();
  });

  test('displays initial value 0', () => {
    const display = screen.getByDisplayValue('0');
    expect(display).toBeInTheDocument();
  });

  test('inputs numbers correctly', () => {
    const button7 = screen.getByText('7');
    fireEvent.click(button7);

    const display = screen.getByDisplayValue('7');
    expect(display).toBeInTheDocument();
  });

  test('handles addition operation', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonAdd = screen.getByText('+');
    const buttonEquals = screen.getByText('=');

    const mockResponse = { result: 13 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    /* fire events that update state */
    fireEvent.click(button1);
    fireEvent.click(button1);
    fireEvent.click(buttonAdd);
    fireEvent.click(button2);
    fireEvent.click(buttonEquals);

    const display = await screen.findByDisplayValue('13');
    expect(display).toBeInTheDocument();
  });

  test('handles square root operation', async () => {
    const button1 = screen.getByText('1');
    const button2 = screen.getByText('2');
    const buttonSqrt = screen.getByText('√');

    const mockResponse = { result: 11 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    fireEvent.click(button1);
    fireEvent.click(button2);
    fireEvent.click(button1);
    fireEvent.click(buttonSqrt);

    const display = await screen.findByDisplayValue('11');
    expect(display).toBeInTheDocument();
  });

  test('handles percentage operation', async () => {
    const button5 = screen.getByText('5');
    const button1 = screen.getByText('1');
    const buttonPercent = screen.getByText('%');
    const buttonEquals = screen.getByText('=');

    const mockResponse = { result: 140 }

    vi.spyOn(globalThis, 'fetch').mockResolvedValue({
      ok: true,
      json: async () => mockResponse,
    } as Response)

    fireEvent.click(button5);
    fireEvent.click(buttonPercent);
    fireEvent.click(button1);
    fireEvent.click(buttonEquals);    

    const display = await screen.findByDisplayValue('140');
    expect(display).toBeInTheDocument();
  });

  test('handles clear operation', () => {
    const button1 = screen.getByText('1');
    const buttonClear = screen.getByText('AC');

    fireEvent.click(button1);
    fireEvent.click(buttonClear);

    const display = screen.getByDisplayValue('0');
    expect(display).toBeInTheDocument();
  });
});