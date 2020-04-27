# Packet Organization
An Ogg Opus stream is organized as follows (see Figure 1 for an
   example).

        Page 0         Pages 1 ... n        Pages (n+1) ...
     +------------+ +---+ +---+ ... +---+ +-----------+ +---------+ +--
     |            | |   | |   |     |   | |           | |         | |
     |+----------+| |+-----------------+| |+-------------------+ +-----
     |||ID Header|| ||  Comment Header || ||Audio Data Packet 1| | ...
     |+----------+| |+-----------------+| |+-------------------+ +-----
     |            | |   | |   |     |   | |           | |         | |
     +------------+ +---+ +---+ ... +---+ +-----------+ +---------+ +--
     ^      ^                           ^
     |      |                           |
     |      |                           Mandatory Page Break
     |      |
     |      ID header is contained on a single page
     |
     'Beginning Of Stream'

    Figure 1: Example Packet Organization for a Logical Ogg Opus Stream
