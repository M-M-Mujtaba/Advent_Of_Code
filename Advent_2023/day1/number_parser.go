package day1

func forwardNumberParser(s []byte) int {
	switch s[0] {
	case 'o':
		switch s[1] {
		case 'n':
			switch s[2] {
			case 'e':
				return 1 + 100
			default:
				return 2
			}
		default:
			return 1
		}
	case 't':
		switch s[1] {
		case 'w':
			switch s[2] {
			case 'o':
				return 2 + 100
			default:
				return 2
			}
		case 'h':
			switch s[2] {
			case 'r':
				switch s[3] {
				case 'e':
					switch s[4] {
					case 'e':
						return 3 + 100
					default:
						return 4

					}
				default:
					return 3
				}
			default:
				return 2

			}
		default:
			return 1
		}
	case 'f':
		switch s[1] {
		case 'o':
			switch s[2] {
			case 'u':
				switch s[3] {
				case 'r':
					return 4 + 100
				default:
					return 3
				}
			default:
				return 2
			}
		case 'i':
			switch s[2] {
			case 'v':
				switch s[3] {
				case 'e':
					return 5 + 100
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}
	case 's':
		switch s[1] {
		case 'i':
			switch s[2] {
			case 'x':
				return 6 + 100
			default:
				return 2

			}
		case 'e':
			switch s[2] {
			case 'v':
				switch s[3] {
				case 'e':
					switch s[4] {
					case 'n':
						return 7 + 100
					default:
						return 4
					}
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1

		}
	case 'e':

		switch s[1] {
		case 'i':
			switch s[2] {
			case 'g':
				switch s[3] {
				case 'h':
					switch s[4] {
					case 't':
						return 8 + 100
					default:
						return 4
					}
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}
	case 'n':
		switch s[1] {
		case 'i':
			switch s[2] {
			case 'n':
				switch s[3] {
				case 'e':
					return 9 + 100
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}

	default:
		return 0
	}

}

func reverseNumberParser(s []byte) int {
	switch s[0] {

	case 'r':
		switch s[1] {
		case 'u':
			switch s[2] {
			case 'o':
				switch s[3] {
				case 'f':
					return 4 + 100
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}
	case 'e':
		switch s[1] {
		case 'v':
			switch s[2] {
			case 'i':
				switch s[3] {
				case 'f':
					return 5 + 100
				default:
					return 3
				}
			default:
				return 2
			}
		case 'n':
			switch s[2] {
			case 'i':
				switch s[3] {
				case 'n':
					return 9 + 100
				default:
					return 3
				}
			case 'o':
				return 1 + 100
			default:
				return 2
			}
		case 'e':
			switch s[2] {
			case 'r':
				switch s[3] {
				case 'h':
					switch s[4] {
					case 't':
						return 3 + 100
					default:
						return 4
					}
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}
	case 'x':
		switch s[1] {
		case 'i':
			switch s[2] {
			case 's':
				return 6 + 100
			default:
				return 2
			}
		default:
			return 1
		}
	case 't':
		switch s[1] {
		case 'h':
			switch s[2] {
			case 'g':
				switch s[3] {
				case 'i':
					switch s[4] {
					case 'e':
						return 8 + 100
					default:
						return 4
					}
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}

	case 'n':
		switch s[1] {
		case 'e':
			switch s[2] {
			case 'v':
				switch s[3] {
				case 'e':
					switch s[4] {
					case 's':
						return 7 + 100
					default:
						return 4

					}
				default:
					return 3
				}
			default:
				return 2
			}
		default:
			return 1
		}
	case 'o':
		switch s[1] {
		case 'w':
			switch s[2] {
			case 't':
				return 2 + 100
			default:
				return 2
			}
		default:
			return 1
		}
	default:
		return 0

	}

}
