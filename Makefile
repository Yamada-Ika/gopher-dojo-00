NAME		:= convert
ifndef WITH_BONUS
	SRCS	:= main.go convert.go
else
	SRCS	:= main_bonus.go
endif

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRCS)

bonus:
	make WITH_BONUS=1

clean:
	rm -rf $(NAME)

fclean: clean

re: fclean all

.PHONY: all clean fclean re